package repository

import (
	"github.com/adityasuryadi/go-shop/services/product/internal/entity"
	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
	db *gorm.DB
	Repository[entity.Product]
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{
		db: db,
	}
}

// Store implements ProductRepository.
func (r *ProductRepositoryImpl) Store(product *entity.Product) (*entity.Product, error) {
	err := r.Repository.Create(r.db, product)
	if err != nil {
		return nil, err
	}
	return product, nil
}
