package repository

import (
	"github.com/adityasuryadi/go-shop/services/auth/internal/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindUserByEmail(db *gorm.DB, email string) (*entity.User, error)
}
