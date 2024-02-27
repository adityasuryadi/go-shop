package model

type CreateProductRequest struct {
	Name        string   `json:"name" validate:"required"`
	Price       int64    `json:"price" validate:"required,gte=1"`
	Stock       int64    `json:"stock" validate:"required,gte=1"`
	Description string   `json:"description" validate:"required"`
	Categories  []string `json:"categories" validate:"required"`
}

type UpdateProductRequest struct {
	Name        string   `json:"name" validate:"required"`
	Price       int64    `json:"price" validate:"required,gte=1"`
	Stock       int64    `json:"stock" validate:"required,gte=1"`
	Description string   `json:"description" validate:"required"`
	Categories  []string `json:"categories" validate:"required"`
}

type Product struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Price       int64  `json:"price"`
	Stock       int64  `json:"stock"`
	Description string `json:"description"`
}

type ProductResponse struct {
	Name        string                     `json:"name"`
	Price       int64                      `json:"price"`
	Stock       int64                      `json:"stock"`
	Description string                     `json:"description"`
	CreatedAt   int64                      `json:"created_at"`
	UpdatedAt   int64                      `json:"updated_at"`
	Categories  []*ProductCategoryResponse `json:"categories"`
}

type ProductCategoryResponse struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type SearchProductRequest struct {
	Page int `json:"page" validate:"required,numeric,min=1"`
	Size int `json:"size" validate:"required,numeric,min=1,max=100"`
}

type FilterProductRequest struct {
	Categories []string `json:"categories"`
}
