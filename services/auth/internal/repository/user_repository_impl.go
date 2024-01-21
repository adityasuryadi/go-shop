package repository

import (
	"github.com/adityasuryadi/go-shop/services/auth/internal/entity"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
}

// FindUserByEmail implements UserRepository.
func (r *UserRepositoryImpl) FindUserByEmail(db *gorm.DB, email string) (*entity.User, error) {
	var user *entity.User
	err := db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func NewUserRespository() UserRepository {
	return &UserRepositoryImpl{}
}
