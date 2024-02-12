package http

import (
	"encoding/json"
	"net/http"

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

func (c *ProductController) InitRoute(Router *chi.Mux) {
	Router.Post("/", c.Create)
}
