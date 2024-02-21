package usecase

import (
	"github.com/adityasuryadi/go-shop/pkg/exception"
	"github.com/adityasuryadi/go-shop/services/product/internal/config"
	"github.com/adityasuryadi/go-shop/services/product/internal/entity"
	"github.com/adityasuryadi/go-shop/services/product/internal/model"
	"github.com/adityasuryadi/go-shop/services/product/internal/model/converter"
	"github.com/adityasuryadi/go-shop/services/product/internal/repository"
)

func NewCategoryUsecase(categoryRepository repository.CategoryRepository, validation *config.Validation) CategoryUsecase {
	return &CategoryUsecaseImpl{
		repository: categoryRepository,
		validation: validation,
	}
}

type CategoryUsecaseImpl struct {
	repository repository.CategoryRepository
	validation *config.Validation
}

func (u *CategoryUsecaseImpl) Create(request *model.CreateCategoryRequest) *exception.CustomError {
	err := u.validation.ValidateRequest(request)
	if err != nil {
		return &exception.CustomError{
			Status: exception.ERRRBADREQUEST,
			Errors: err,
		}
	}
	category := &entity.Category{
		Name: request.Name,
	}

	err = u.repository.Create(category)
	if err != nil {
		return &exception.CustomError{
			Status: exception.ERRBUSSINESS,
			Errors: err,
		}
	}
	return nil
}

func (u *CategoryUsecaseImpl) Search(request *model.SearchCategoryRequest) ([]*model.CategoryResponse, int64, *exception.CustomError) {
	err := u.validation.ValidateRequest(request)
	if err != nil {
		return nil, 0, &exception.CustomError{
			Status: exception.ERRRBADREQUEST,
			Errors: err,
		}
	}

	result, total, err := u.repository.Search(request)
	responses := make([]*model.CategoryResponse, len(result))
	for i, v := range result {
		responses[i] = converter.CategoryToResponse(&v)
	}

	return responses, total, nil
}
