package usecase

import (
	"fmt"

	"github.com/adityasuryadi/go-shop/pkg/logger"
	"github.com/adityasuryadi/go-shop/pkg/util"
	"github.com/adityasuryadi/go-shop/services/auth/internal/config"
	"github.com/adityasuryadi/go-shop/services/auth/internal/model"
	"github.com/adityasuryadi/go-shop/services/auth/internal/repository"
	"gorm.io/gorm"
)

type AuthUsecaseImpl struct {
	db             *gorm.DB
	userRepository repository.UserRepository
	jwtConfig      *config.JwtConfig
	logger         *logger.Logger
	validation     *config.Validation
}

func NewAuthUsecase(userRepo repository.UserRepository, db *gorm.DB, jwtConfig *config.JwtConfig, logger *logger.Logger, validation *config.Validation) AuthUsecase {
	return &AuthUsecaseImpl{
		db:             db,
		userRepository: userRepo,
		jwtConfig:      jwtConfig,
		logger:         logger,
		validation:     validation,
	}
}

func (u *AuthUsecaseImpl) Login(request *model.LoginRequest) (string, error) {
	err := u.validation.ValidateRequest(request)
	if err != nil {
		return "", err
	}

	db := u.db.Begin()
	email := request.Email
	user, err := u.userRepository.FindUserByEmail(db, email)
	if err != nil {
		u.logger.Errorf("failed to find user ", err)
		return "", nil
	}
	fmt.Println(user)
	err = util.ComparePassword(user.Password, request.Password)
	if err != nil {
		u.logger.Errorf("failed to compare password ", err)
		return "", nil
	}
	accessToken, err := u.jwtConfig.ClaimAccessToken(email)
	if err != nil {
		u.logger.Errorf("failed to claim token ", err)
		return "", nil
	}
	db.Commit()
	return accessToken, nil
}
