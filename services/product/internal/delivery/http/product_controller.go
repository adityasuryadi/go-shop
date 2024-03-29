package http

import (
	"encoding/json"
	"math"
	"net/http"
	"strconv"

	"github.com/adityasuryadi/go-shop/pkg/exception"
	"github.com/adityasuryadi/go-shop/pkg/logger"
	"github.com/adityasuryadi/go-shop/services/product/internal/config"
	"github.com/adityasuryadi/go-shop/services/product/internal/model"
	"github.com/adityasuryadi/go-shop/services/product/internal/usecase"
	"github.com/go-chi/chi/v5"
)

type ProductController struct {
	Router  *chi.Mux
	Usecase usecase.ProductUsecase
	logger  *logger.Logger
}

func NewProductController(r *chi.Mux, productUsecase usecase.ProductUsecase, log *logger.Logger) ProductController {
	return ProductController{
		Router:  r,
		Usecase: productUsecase,
		logger:  log,
	}
}

func (c *ProductController) Create(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	request := new(model.CreateProductRequest)
	if err := decoder.Decode(&request); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Bad Request")
		return
	}

	validation := new(config.Validation)
	result, err := c.Usecase.Create(request)
	if err != nil && err.Status == exception.ERRRBADREQUEST {
		errValidation := validation.ErrorJson(err.Errors)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse[any]{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
			Error:  errValidation,
		})
		return
	}
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Bad Request")
		return
	}

	response := model.WebResponse[*model.ProductResponse]{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   result,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (c *ProductController) FindProductById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	result, err := c.Usecase.FindById(id)

	if err != nil && err.Status == exception.ERRNOTFOUND {
		errResponse := model.ErrorResponse[any]{
			Code:   http.StatusNotFound,
			Status: "NOT_FOUND",
			Error:  err.Errors.Error(),
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	if err != nil {
		errResponse := model.ErrorResponse[any]{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
			Error:  err.Errors,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	response := model.WebResponse[*model.ProductResponse]{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   result,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

}

// search product
func (c *ProductController) Search(w http.ResponseWriter, r *http.Request) {
	// TODO
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		page = 1
	}
	size, err := strconv.Atoi(r.URL.Query().Get("size"))
	if err != nil {
		size = 10
	}
	request := &model.SearchProductRequest{
		Page: page,
		Size: size,
	}
	validation := new(config.Validation)
	result, total, customErr := c.Usecase.Search(request)

	if customErr != nil && customErr.Status == exception.ERRRBADREQUEST {
		errValidation := validation.ErrorJson(customErr.Errors)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse[any]{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
			Error:  errValidation,
		})
		return
	}

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Bad Request")
		return
	}

	response := model.WebResponse[[]*model.ProductResponse]{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   result,
		Paging: &model.PageMetadata{
			Page:      page,
			Size:      size,
			TotalItem: total,
			TotalPage: int64(math.Ceil(float64(total) / float64(size))),
		},
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (c *ProductController) FilterProduct(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	request := new(model.FilterProductRequest)
	if err := decoder.Decode(&request); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Bad Request")
		return
	}
	response, err := c.Usecase.FilterProduct(request)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Bad Request")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (c *ProductController) Update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	decoder := json.NewDecoder(r.Body)
	request := new(model.UpdateProductRequest)
	if err := decoder.Decode(&request); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Bad Request")
		return
	}

	validation := new(config.Validation)
	result, err := c.Usecase.Update(id, request)
	if err != nil && err.Status == exception.ERRRBADREQUEST {
		errValidation := validation.ErrorJson(err.Errors)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse[any]{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
			Error:  errValidation,
		})
		return
	}

	if err != nil && err.Status == exception.ERRNOTFOUND {
		errResponse := model.ErrorResponse[any]{
			Code:   http.StatusNotFound,
			Status: "NOT_FOUND",
			Error:  err.Errors.Error(),
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	if err != nil && err.Status == exception.ERRBUSSINESS {
		errResponse := model.ErrorResponse[any]{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL_SERVER_ERROR",
			Error:  err.Errors,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	response := model.WebResponse[*model.ProductResponse]{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   result,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (c *ProductController) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := c.Usecase.Delete(id)

	if err != nil && err.Status == exception.ERRNOTFOUND {
		errResponse := model.ErrorResponse[any]{
			Code:   http.StatusNotFound,
			Status: "NOT_FOUND",
			Error:  err.Errors.Error(),
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	if err != nil {
		errResponse := model.ErrorResponse[any]{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
			Error:  err.Errors,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	response := model.WebResponse[*model.ProductResponse]{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (c *ProductController) InitRoute(Router *chi.Mux) {
	Router.Post("/product", c.Create)
	Router.Get("/product/{id}", c.FindProductById)
	Router.Get("/product/search", c.Search)
	Router.Put("/product/{id}", c.Update)
	Router.Delete("/product/{id}", c.Delete)
	Router.Post("/product/filter", c.FilterProduct)
}
