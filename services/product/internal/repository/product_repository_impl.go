package repository

import (
	"github.com/adityasuryadi/go-shop/services/product/internal/entity"
	"github.com/adityasuryadi/go-shop/services/product/internal/model"
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

// search product
func (r *ProductRepositoryImpl) Search(request *model.SearchProductRequest) ([]entity.Product, int64, error) {
	var products []entity.Product
	err := r.db.Scopes(r.FilterOrder(request)).Offset((request.Page - 1) * request.Size).Limit(request.Size).Find(&products).Error
	if err != nil {
		return nil, 0, err
	}
	var total int64 = 0
	if err := r.db.Model(&entity.Product{}).Scopes(r.FilterOrder(request)).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	return products, total, nil
}

func (r *ProductRepositoryImpl) FilterOrder(request *model.SearchProductRequest) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		return tx
	}
}

// Store implements ProductRepository.
func (r *ProductRepositoryImpl) Store(tx *gorm.DB, product *entity.Product) (*entity.Product, error) {
	// err := r.Repository.Create(r.db, product)
	err := tx.Omit("Categories.*").Create(&product).Debug().Error
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r *ProductRepositoryImpl) AssignCategory(tx *gorm.DB, product *entity.Product, categories []*entity.Category) error {
	tx.Model(&product).Association("Categories").Clear()
	tx.Model(&product).Association("Categories").Append(categories)
	return tx.Error
}

func (r *ProductRepositoryImpl) FindById(id string) (*entity.Product, error) {
	product := new(entity.Product)
	err := r.Repository.FindById(r.db, product, id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r *ProductRepositoryImpl) Update(product *entity.Product) (*entity.Product, error) {
	err := r.Repository.Update(r.db, product)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r *ProductRepositoryImpl) Delete(product *entity.Product) error {
	return r.Repository.Delete(r.db, product)
}
