package usecase

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/adityasuryadi/go-shop/pkg/exception"
	"github.com/adityasuryadi/go-shop/pkg/logger"
	"github.com/adityasuryadi/go-shop/pkg/util"
	"github.com/adityasuryadi/go-shop/services/auth/internal/config"
	"github.com/adityasuryadi/go-shop/services/auth/internal/model"
	"github.com/adityasuryadi/go-shop/services/auth/internal/repository"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type AuthUsecaseImpl struct {
	db             *gorm.DB
	userRepository repository.UserRepository
	jwtConfig      *config.JwtConfig
	logger         *logger.Logger
	validation     *config.Validation
	redisClient    *redis.Client
}

func NewAuthUsecase(userRepo repository.UserRepository, db *gorm.DB, jwtConfig *config.JwtConfig, logger *logger.Logger, validation *config.Validation, redisClient *redis.Client) AuthUsecase {
	return &AuthUsecaseImpl{
		db:             db,
		userRepository: userRepo,
		jwtConfig:      jwtConfig,
		logger:         logger,
		validation:     validation,
		redisClient:    redisClient,
	}
}

func (u *AuthUsecaseImpl) Login(request *model.LoginRequest) (*model.LoginResponse, *exception.CustomError) {
	ctx := context.Background()
	err := u.validation.ValidateRequest(request)
	if err != nil {
		return nil, &exception.CustomError{
			Status: exception.ERRRBADREQUEST,
			Errors: err,
		}
	}

	db := u.db.Begin()
	email := request.Email
	user, err := u.userRepository.FindUserByEmail(db, email)
	if err != nil {
		u.logger.Errorf("failed to find user ", err)
		return nil, &exception.CustomError{
			Status: exception.ERRNOTFOUND,
			Errors: errors.New("user not found"),
		}
	}

	err = util.ComparePassword(user.Password, request.Password)
	if err != nil {
		u.logger.Errorf("failed to compare password ", err)
		return nil, &exception.CustomError{
			Status: exception.ERRNOTFOUND,
			Errors: errors.New("wrong email or password"),
		}
	}
	accessToken, err := u.jwtConfig.ClaimAccessToken(email)
	if err != nil {
		u.logger.Errorf("failed to claim token ", err)
		return nil, &exception.CustomError{
			Status: exception.ERRNOTFOUND,
			Errors: errors.New("failed claim token"),
		}
	}

	refreshToken, err := u.jwtConfig.ClaimRefreshToken(email)
	db.Commit()
	u.redisClient.Set(ctx, "refresh_token:"+email+":"+refreshToken, refreshToken, time.Duration(time.Hour*7))
	response := &model.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	return response, nil
}

func (u *AuthUsecaseImpl) RefreshToken(refreshToken string) (*model.LoginResponse, *exception.CustomError) {
	claims := u.jwtConfig.DecodeTokenString(refreshToken)
	email := claims["email"].(string)
	// get token from redis
	ctx := context.Background()
	redisKey := "refresh_token:" + email + ":" + refreshToken
	value := u.redisClient.Get(ctx, redisKey)
	if value.Val() == "" {
		return nil, &exception.CustomError{
			Errors: errors.New("unauthorized"),
			Status: exception.ERRAUTHORIZED,
		}
	}

	accessToken, err := u.jwtConfig.ClaimAccessToken(email)
	if err != nil {
		u.logger.Errorf("failed to claim token ", err.Error())
		return nil, &exception.CustomError{
			Status: exception.ERRSYSTEM,
			Errors: err,
		}
	}
	response := &model.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	return response, nil
}

func (u *AuthUsecaseImpl) Logout(token string) error {
	token = strings.Replace(token, "Bearer ", "", 1)
	claims := u.jwtConfig.DecodeTokenString(token)
	email := claims["email"].(string)
	redisKey := "refresh_token:" + email + ":" + token
	u.redisClient.Del(context.Background(), redisKey)

	return nil
}
