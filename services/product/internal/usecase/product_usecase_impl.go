package usecase

import (
	"errors"
	"fmt"

	"github.com/adityasuryadi/go-shop/pkg/exception"
	"github.com/adityasuryadi/go-shop/services/product/internal/config"
	"github.com/adityasuryadi/go-shop/services/product/internal/entity"
	"github.com/adityasuryadi/go-shop/services/product/internal/model"
	"github.com/adityasuryadi/go-shop/services/product/internal/model/converter"
	"github.com/adityasuryadi/go-shop/services/product/internal/repository"
	"gorm.io/gorm"
)

type ProductUsecaseImpl struct {
	db          *gorm.DB
	productRepo repository.ProductRepository
	validation  *config.Validation
}

func NewProductUsecase(db *gorm.DB, productRepository repository.ProductRepository, validation *config.Validation) ProductUsecase {
	return &ProductUsecaseImpl{
		db:          db,
		productRepo: productRepository,
		validation:  validation,
	}
}

// search product
func (u *ProductUsecaseImpl) Search(request *model.SearchProductRequest) ([]*model.ProductResponse, int64, *exception.CustomError) {
	err := u.validation.ValidateRequest(request)
	if err != nil {
		return nil, 0, &exception.CustomError{
			Status: exception.ERRRBADREQUEST,
			Errors: err,
		}
	}

	result, total, err := u.productRepo.Search(request)
	responses := make([]*model.ProductResponse, len(result))
	for i, v := range result {
		responses[i] = converter.ProductToResponse(&v)
	}

	return responses, total, nil
}

// FindById implements ProductUsecase.
func (u *ProductUsecaseImpl) FindById(id string) (*model.ProductResponse, *exception.CustomError) {
	if id == "" {
		return nil, &exception.CustomError{
			Status: exception.ERRRBADREQUEST,
			Errors: errors.New("id required"),
		}
	}

	result, err := u.productRepo.FindById(id)
	fmt.Println("result", result)
	fmt.Println("err ", err)
	if err != nil {
		return nil, &exception.CustomError{
			Status: exception.ERRNOTFOUND,
			Errors: errors.New("product not found"),
		}
	}

	response := converter.ProductToResponse(result)
	fmt.Println(response)
	return response, nil
}

// Create implements ProductUsecase.
func (u *ProductUsecaseImpl) Create(request *model.CreateProductRequest) (*model.ProductResponse, *exception.CustomError) {
	product := &entity.Product{
		Name:        request.Name,
		Stock:       request.Stock,
		Description: request.Description,
		Price:       request.Price,
	}

	// validasi request
	err := u.validation.ValidateRequest(request)
	if err != nil {
		return nil, &exception.CustomError{
			Status: exception.ERRRBADREQUEST,
			Errors: err,
		}
	}

	result, err := u.productRepo.Store(product)
	if err != nil {
		return nil, &exception.CustomError{
			Status: exception.ERRBUSSINESS,
			Errors: err,
		}
	}

	response := &model.ProductResponse{
		Name:        result.Name,
		Price:       result.Price,
		Stock:       result.Stock,
		Description: result.Description,
	}
	return response, nil
}

// update product
func (u *ProductUsecaseImpl) Update(id string, request *model.UpdateProductRequest) (*model.ProductResponse, *exception.CustomError) {
	// validasi product
	if id == "" {
		return nil, &exception.CustomError{
			Status: exception.ERRRBADREQUEST,
			Errors: errors.New("id required"),
		}
	}

	err := u.validation.ValidateRequest(request)
	if err != nil {
		return nil, &exception.CustomError{
			Status: exception.ERRRBADREQUEST,
			Errors: err,
		}
	}

	// find product
	product, err := u.productRepo.FindById(id)
	if err != nil {
		return nil, &exception.CustomError{
			Status: exception.ERRNOTFOUND,
			Errors: errors.New("product not found"),
		}
	}

	product.Name = request.Name
	product.Price = request.Price
	product.Stock = request.Stock
	product.Description = request.Description

	product, err = u.productRepo.Update(product)
	if err != nil {
		return nil, &exception.CustomError{
			Status: exception.ERRBUSSINESS,
			Errors: err,
		}
	}

	response := converter.ProductToResponse(product)

	return response, nil
}

func (u *ProductUsecaseImpl) Delete(id string) *exception.CustomError {
	// validasi product
	if id == "" {
		return &exception.CustomError{
			Status: exception.ERRRBADREQUEST,
			Errors: errors.New("id required"),
		}
	}

	// find product
	product, err := u.productRepo.FindById(id)
	if err != nil {
		return &exception.CustomError{
			Status: exception.ERRNOTFOUND,
			Errors: errors.New("product not found"),
		}
	}

	err = u.productRepo.Delete(product)
	if err != nil {
		return &exception.CustomError{
			Status: exception.ERRBUSSINESS,
			Errors: err,
		}
	}
	return nil
}
