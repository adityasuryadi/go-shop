package model

type CreateProductRequest struct {
	Name        string `json:"name" validate:"required"`
	Price       int64  `json:"price" validate:"required,gte=1"`
	Stock       int64  `json:"stock" validate:"required,gte=1"`
	Description string `json:"description" validate:"required"`
}

type UpdateProductRequest struct {
	Name        string `json:"name"`
	Price       int64  `json:"price"`
	Stock       int64  `json:"stock"`
	Description string `json:"description"`
}

type Product struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Price       int64  `json:"price"`
	Stock       int64  `json:"stock"`
	Description string `json:"description"`
}

type ProductResponse struct {
	Name        string `json:"name"`
	Price       int64  `json:"price"`
	Stock       int64  `json:"stock"`
	Description string `json:"description"`
}
