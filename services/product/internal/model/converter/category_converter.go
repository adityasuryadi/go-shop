package converter

import (
	"github.com/adityasuryadi/go-shop/services/product/internal/entity"
	"github.com/adityasuryadi/go-shop/services/product/internal/model"
)

func CategoryToResponse(category *entity.Category) *model.CategoryResponse {
	return &model.CategoryResponse{
		Id:        category.Id,
		Name:      category.Name,
		IsActive:  category.IsActive,
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
	}
}
