package model

import "github.com/google/uuid"

type Category struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	IsActive  bool      `json:"is_active"`
	CreatedAt int64     `json:"created_at"`
	UpdatedAt int64     `json:"updated_at"`
}

type CreateCategoryRequest struct {
	Name string `json:"name" validate:"required"`
}

type UpdateCategoryRequest struct {
	Name string `json:"name" validate:"required"`
}

type CategoryResponse struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	IsActive  int8      `json:"is_active"`
	CreatedAt int64     `json:"created_at"`
	UpdatedAt int64     `json:"updated_at"`
}

type SearchCategoryRequest struct {
	Page int `json:"page" validate:"required,numeric,min=1"`
	Size int `json:"size" validate:"required,numeric,min=1,max=100"`
}
