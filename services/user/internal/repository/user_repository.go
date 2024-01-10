package repository

import (
	"github.com/adityasuryadi/go-shop/services/user/internal/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	Insert(db *gorm.DB, user *entity.User) (*entity.User, error)
}
