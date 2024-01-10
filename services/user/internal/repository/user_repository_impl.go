package repository

import (
	"github.com/adityasuryadi/go-shop/services/user/internal/entity"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

// Insert implements UserRepository.
func (*UserRepositoryImpl) Insert(db *gorm.DB, user *entity.User) (*entity.User, error) {
	err := db.Create(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
