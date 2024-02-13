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

func NewProductUsecase(db *gorm.DB, productRepository repository.ProductRepository, validation *config.Validation) ProductUsecase {
	return &ProductUsecaseImpl{
		db:          db,
		productRepo: productRepository,
		validation:  validation,
	}
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
