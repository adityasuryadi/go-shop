package repository

import (
	"github.com/adityasuryadi/go-shop/services/user/internal/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(db *gorm.DB, entity *entity.User) (*entity.User, error)
	FindById(db *gorm.DB, entity *entity.User, id any) error
}
