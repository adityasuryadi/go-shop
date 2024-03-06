package usecase

import (
	"context"

	"github.com/adityasuryadi/go-shop/pkg/exception"
	"github.com/adityasuryadi/go-shop/services/user/internal/model"
)

type UserUsecase interface {
	Insert(context.Context, *model.CreateUserRequest) (*model.UserResponse, *exception.CustomError)
	FindById(id string) (*model.UserResponse, error)
}
