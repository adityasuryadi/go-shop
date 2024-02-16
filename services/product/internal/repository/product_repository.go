package repository

import (
	"github.com/adityasuryadi/go-shop/services/product/internal/entity"
	"github.com/adityasuryadi/go-shop/services/product/internal/model"
)

type ProductRepository interface {
	Store(product *entity.Product) (*entity.Product, error)
	FindById(id string) (*entity.Product, error)
	Search(request *model.SearchProductRequest) ([]entity.Product, int64, error)
	Update(product *entity.Product) (*entity.Product, error)
	// FindAll() ([]entity.Product, error)
}
