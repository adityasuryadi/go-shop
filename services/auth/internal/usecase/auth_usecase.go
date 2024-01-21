package usecase

import "github.com/adityasuryadi/go-shop/services/auth/internal/model"

type AuthUsecase interface {
	Login(request *model.LoginRequest) (string, error)
}
