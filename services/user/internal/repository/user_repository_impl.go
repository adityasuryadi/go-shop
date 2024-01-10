package repository

import (
	"github.com/adityasuryadi/go-shop/pkg/logger"
	"github.com/adityasuryadi/go-shop/services/user/internal/entity"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	log *logger.Logger
	Repository[entity.User]
}

func NewUserRepository(log *logger.Logger) UserRepository {
	return &UserRepositoryImpl{
		log: log,
	}
}

// Insert implements UserRepository.
func (r *UserRepositoryImpl) Insert(db *gorm.DB, user *entity.User) (*entity.User, error) {
	err := db.Create(user).Error
	if err != nil {
		r.log.Errorf("failed to create user ", err)
		return nil, err
	}
	return user, nil
}
