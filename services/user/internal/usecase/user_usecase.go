package usecase

import (
	"context"

	"github.com/adityasuryadi/go-shop/services/user/internal/model"
)

type UserUsecase interface {
	Insert(context.Context, *model.CreateUserRequest) (*model.UserResponse, error)
	FindById(id string) (*model.UserResponse, error)
}
