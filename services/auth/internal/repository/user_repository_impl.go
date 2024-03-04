package repository

import (
	"time"

	"github.com/adityasuryadi/go-shop/pkg/logger"
	"github.com/adityasuryadi/go-shop/services/auth/internal/entity"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	log *logger.Logger
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

func (r *UserRepositoryImpl) Insert(db *gorm.DB, user *entity.User) (*entity.User, error) {
	err := db.Create(user).Error
	if err != nil {
		r.log.Errorf("failed to create user ", err)
		return nil, err
	}
	return user, nil
}

func (r *UserRepositoryImpl) VerifyUser(db *gorm.DB, email string) error {
	err := db.Where("email = ?", email).Update("verified_at", time.Now().Unix()).Error
	if err != nil {
		return err
	}
	return nil
}

func NewUserRespository(logger *logger.Logger) UserRepository {
	return &UserRepositoryImpl{
		log: logger,
	}
}
