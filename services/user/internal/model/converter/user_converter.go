package converter

import (
	"github.com/adityasuryadi/go-shop/services/user/internal/entity"
	"github.com/adityasuryadi/go-shop/services/user/internal/model"
)

func UserToResponse(user *entity.User) *model.UserResponse {
	return &model.UserResponse{
		FirstName: user.FirtsName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     user.Phone,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
