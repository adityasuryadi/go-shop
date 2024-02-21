package repository

import (
	"github.com/adityasuryadi/go-shop/services/product/internal/entity"
	"github.com/adityasuryadi/go-shop/services/product/internal/model"
)

type CategoryRepository interface {
	Create(category *entity.Category) error
	Search(request *model.SearchCategoryRequest) ([]entity.Category, int64, error)
}
