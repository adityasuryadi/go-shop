package usecase

import (
	"context"

	"github.com/adityasuryadi/go-shop/pkg/exception"
	"github.com/adityasuryadi/go-shop/pkg/logger"
	"github.com/adityasuryadi/go-shop/services/user/internal/config"
	"github.com/adityasuryadi/go-shop/services/user/internal/entity"
	"github.com/adityasuryadi/go-shop/services/user/internal/gateway/messaging"
	"github.com/adityasuryadi/go-shop/services/user/internal/model"
	"github.com/adityasuryadi/go-shop/services/user/internal/model/converter"
	"github.com/adityasuryadi/go-shop/services/user/internal/repository"
	hash "github.com/adityasuryadi/go-shop/services/user/internal/utils"
	"github.com/rabbitmq/amqp091-go"

	"gorm.io/gorm"
)

type UserUsecaseImpl struct {
	db         *gorm.DB
	repository repository.UserRepository
	channel    *amqp091.Channel
	validation *config.Validation
}

// Insert implements UserUsecase.
func (u *UserUsecaseImpl) Insert(ctx context.Context, request *model.CreateUserRequest) (*model.UserResponse, *exception.CustomError) {
	tx := u.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	err := u.validation.ValidateRequest(request)
	if err != nil {
		return nil, &exception.CustomError{
			Status: exception.ERRRBADREQUEST,
			Errors: err,
		}
	}
	apilog := logger.NewLogger()

	hashPassword, err := hash.HashPassword([]byte(request.Password))
	if err != nil {
		apilog.Errorf("failed to create user ", err)
		return nil, &exception.CustomError{
			Status: exception.ERRRBADREQUEST,
			Errors: err,
		}
	}
	user := &entity.User{
		FirtsName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
		Phone:     request.Phone,
		Password:  hashPassword,
	}
	result, err := u.repository.Create(tx, user)
	if err != nil {
		apilog.Errorf("failed to create user ", err)
		return nil, &exception.CustomError{
			Status: exception.ERRDOMAIN,
			Errors: err,
		}
	}

	response := converter.UserToResponse(result)

	if err := tx.Commit().Error; err != nil {
		apilog.Errorf("failed to create user ", err)
		return nil, &exception.CustomError{
			Status: exception.ERRDOMAIN,
			Errors: err,
		}
	}

	userProducerConfig := &messaging.ProducerConfig{
		Exchange:   "user.created",
		QueueName:  "user.create",
		RoutingKey: "create",
	}
	userProducer := messaging.NewUserProducer(u.channel, userProducerConfig, apilog)
	userEvent := &model.CreateUserEvent{
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
		Phone:     request.Phone,
		Password:  hashPassword,
	}

	// defer u.channel.Close()
	userProducer.SetupExchangeAndQueuePublisher()
	userProducer.Publish(userEvent)
	return response, nil
}

func (u *UserUsecaseImpl) FindById(id string) (*model.UserResponse, error) {
	tx := u.db
	user := new(entity.User)
	err := u.repository.FindById(tx, user, id)
	if err != nil {
		return nil, err
	}
	response := converter.UserToResponse(user)
	return response, nil
}

func NewUserUsecase(db *gorm.DB, userRepo repository.UserRepository, channel *amqp091.Channel, validate *config.Validation) UserUsecase {
	return &UserUsecaseImpl{
		db:         db,
		repository: userRepo,
		channel:    channel,
		validation: validate,
	}
}
