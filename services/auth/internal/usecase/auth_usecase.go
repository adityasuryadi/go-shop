package usecase

import (
	"github.com/adityasuryadi/go-shop/pkg/exception"
	"github.com/adityasuryadi/go-shop/services/auth/internal/model"
)

type AuthUsecase interface {
	Login(request *model.LoginRequest) (*model.LoginResponse, *exception.CustomError)
	RefreshToken(refreshToken string) (*model.LoginResponse, *exception.CustomError)
	Logout(token string) error
}
