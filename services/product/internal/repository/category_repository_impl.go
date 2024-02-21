package repository

import (
	"github.com/adityasuryadi/go-shop/services/product/internal/entity"
	"github.com/adityasuryadi/go-shop/services/product/internal/model"
	"gorm.io/gorm"
)

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &CategoryRepositoryImpl{
		db: db,
	}
}

type CategoryRepositoryImpl struct {
	repository Repository[entity.Category]
	db         *gorm.DB
}

func (r *CategoryRepositoryImpl) Create(category *entity.Category) error {
	err := r.repository.Create(r.db, category)
	if err != nil {
		return err
	}
	return nil
}

// search product
func (r *CategoryRepositoryImpl) Search(request *model.SearchCategoryRequest) ([]entity.Category, int64, error) {
	var categories []entity.Category
	err := r.db.Scopes(r.FilterCategory(request)).Offset((request.Page - 1) * request.Size).Limit(request.Size).Find(&categories).Error
	if err != nil {
		return nil, 0, err
	}
	var total int64 = 0
	if err := r.db.Model(&entity.Category{}).Scopes(r.FilterCategory(request)).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	return categories, total, nil
}

func (r *CategoryRepositoryImpl) FilterCategory(request *model.SearchCategoryRequest) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		return tx
	}
}
