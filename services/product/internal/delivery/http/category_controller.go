package http

import (
	"encoding/json"
	"math"
	"net/http"
	"strconv"

	"github.com/adityasuryadi/go-shop/pkg/exception"
	"github.com/adityasuryadi/go-shop/services/product/internal/config"
	"github.com/adityasuryadi/go-shop/services/product/internal/model"
	"github.com/adityasuryadi/go-shop/services/product/internal/usecase"
	"github.com/go-chi/chi/v5"
)

type CategoryController struct {
	Router  *chi.Mux
	Usecase usecase.CategoryUsecase
}

func NewCategoryController(router *chi.Mux, usecase usecase.CategoryUsecase) *CategoryController {
	return &CategoryController{
		Router:  router,
		Usecase: usecase,
	}
}

func (c *CategoryController) Create(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	request := new(model.CreateCategoryRequest)
	if err := decoder.Decode(&request); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Bad Request")
		return
	}

	validation := new(config.Validation)
	err := c.Usecase.Create(request)
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

	if err != nil && err.Status == exception.ERRBUSSINESS {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse[any]{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL_SERVER_ERROR",
			Error:  err.Errors,
		})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.WebResponse[any]{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	})

}

func (c *CategoryController) Search(w http.ResponseWriter, r *http.Request) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		page = 1
	}

	size, err := strconv.Atoi(r.URL.Query().Get("size"))
	if err != nil {
		size = 10
	}

	request := &model.SearchCategoryRequest{
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

	response := model.WebResponse[[]*model.CategoryResponse]{
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

func (c *CategoryController) InitRoute(Router *chi.Mux) {
	Router.Post("/category", c.Create)
	Router.Get("/category/search", c.Search)
}
