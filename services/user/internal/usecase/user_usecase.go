package usecase

import (
	"context"

	"github.com/adityasuryadi/go-shop/services/user/internal/model"
)

type UserUsecase interface {
	Insert(context.Context, *model.CreateUserRequest)
}
