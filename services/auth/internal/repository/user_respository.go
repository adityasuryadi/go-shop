package repository

import (
	"github.com/adityasuryadi/go-shop/services/auth/internal/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	Insert(db *gorm.DB, user *entity.User) (*entity.User, error)
	FindUserByEmail(db *gorm.DB, email string) (*entity.User, error)
	VerifyUser(db *gorm.DB, email string) error
}
