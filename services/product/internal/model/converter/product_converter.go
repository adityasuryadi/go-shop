package converter

import (
	"github.com/adityasuryadi/go-shop/services/product/internal/entity"
	"github.com/adityasuryadi/go-shop/services/product/internal/model"
)

func ProductToResponse(product *entity.Product) *model.ProductResponse {
	return &model.ProductResponse{
		Name:        product.Name,
		Price:       product.Price,
		Stock:       product.Stock,
		Description: product.Description,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
		Categories:  ProductCategoriesToResponse(product.Categories),
	}
}

func ProductCategoriesToResponse(categories []entity.Category) []*model.ProductCategoryResponse {
	responses := make([]*model.ProductCategoryResponse, len(categories))
	for i, v := range categories {
		responses[i] = &model.ProductCategoryResponse{
			Id:   v.Id.String(),
			Name: v.Name,
		}
	}
	return responses

}
