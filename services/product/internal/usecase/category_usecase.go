package usecase

import (
	"github.com/adityasuryadi/go-shop/pkg/exception"
	"github.com/adityasuryadi/go-shop/services/product/internal/model"
)

type CategoryUsecase interface {
	Create(request *model.CreateCategoryRequest) *exception.CustomError
	Search(request *model.SearchCategoryRequest) ([]*model.CategoryResponse, int64, *exception.CustomError)
}
