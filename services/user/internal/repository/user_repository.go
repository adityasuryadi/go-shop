package repository

import (
	"github.com/adityasuryadi/go-shop/services/user/internal/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(db *gorm.DB, entity *entity.User) (*entity.User, error)
	Update(db *gorm.DB, entity *entity.User) error
	Delete(db *gorm.DB, entity *entity.User) error
	CountById(db *gorm.DB, id any) (int64, error)
	FindById(db *gorm.DB, entity *entity.User, id any) error
}
