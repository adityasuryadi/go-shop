package usecase

import (
	"github.com/adityasuryadi/go-shop/pkg/exception"
	"github.com/adityasuryadi/go-shop/services/product/internal/model"
)

type ProductUsecase interface {
	Create(request *model.CreateProductRequest) (response *model.ProductResponse, err *exception.CustomError)
	FindById(id string) (response *model.ProductResponse, err *exception.CustomError)
	Search(request *model.SearchProductRequest) (response []*model.ProductResponse, total int64, err *exception.CustomError)
	Update(id string, request *model.UpdateProductRequest) (response *model.ProductResponse, err *exception.CustomError)
}
