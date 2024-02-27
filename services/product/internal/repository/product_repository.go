package repository

import (
	"github.com/adityasuryadi/go-shop/services/product/internal/entity"
	"github.com/adityasuryadi/go-shop/services/product/internal/model"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Store(tx *gorm.DB, product *entity.Product) (*entity.Product, error)
	FindById(id string) (*entity.Product, error)
	Search(request *model.SearchProductRequest) ([]entity.Product, int64, error)
	Update(tx *gorm.DB, product *entity.Product) (*entity.Product, error)
	Delete(product *entity.Product) error
	AssignCategory(tx *gorm.DB, product *entity.Product, categories []*entity.Category) error
	FilterProduct(request *model.FilterProductRequest) ([]entity.Product, error)
	// FindAll() ([]entity.Product, error)
}
