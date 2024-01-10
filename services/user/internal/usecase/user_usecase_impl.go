package usecase

import (
	"context"

	"github.com/adityasuryadi/go-shop/services/user/internal/model"

	"gorm.io/gorm"
)

type UserUsecaseImpl struct {
	tx *gorm.DB
}

// Insert implements UserUsecase.
func (UserUsecaseImpl) Insert(context.Context, *model.CreateUserRequest) {
	panic("unimplemented")
}

func NewUserUsecase(db *gorm.DB) UserUsecase {
	return &UserUsecaseImpl{
		tx: db,
	}
}
