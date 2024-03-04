package usecase

import (
	"context"

	"github.com/adityasuryadi/go-shop/pkg/logger"
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
}

// Insert implements UserUsecase.
func (u *UserUsecaseImpl) Insert(ctx context.Context, request *model.CreateUserRequest) (*model.UserResponse, error) {
	tx := u.db.WithContext(ctx).Begin()
	defer tx.Rollback()
	apilog := logger.NewLogger()

	hashPassword, err := hash.HashPassword([]byte(request.Password))
	if err != nil {
		apilog.Errorf("failed to create user ", err)
		return nil, err
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
		return nil, err
	}

	response := converter.UserToResponse(result)

	if err := tx.Commit().Error; err != nil {
		apilog.Errorf("failed to create user ", err)
		return nil, err
	}

	userProducerConfig := &messaging.ProducerConfig{
		Exchange:   "user.created",
		QueueName:  "user.create",
		RoutingKey: "user.create",
	}
	userProducer := messaging.NewUserProducer(u.channel, userProducerConfig, apilog)
	userEvent := &model.CreateUserEvent{
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
		Phone:     request.Phone,
		Password:  hashPassword,
	}

	defer u.channel.Close()
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

func NewUserUsecase(db *gorm.DB, userRepo repository.UserRepository, channel *amqp091.Channel) UserUsecase {
	return &UserUsecaseImpl{
		db:         db,
		repository: userRepo,
		channel:    channel,
	}
}
