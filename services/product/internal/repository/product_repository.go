package repository

import (
	"github.com/adityasuryadi/go-shop/services/product/internal/entity"
)

type ProductRepository interface {
	Store(product *entity.Product) (*entity.Product, error)
	FindById(id string) (*entity.Product, error)
	// FindAll() ([]entity.Product, error)
}