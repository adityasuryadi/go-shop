package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/adityasuryadi/go-shop/pkg/logger"
	"github.com/adityasuryadi/go-shop/services/auth/internal/config"
	"github.com/adityasuryadi/go-shop/services/auth/internal/delivery/messaging"
	"github.com/adityasuryadi/go-shop/services/auth/internal/repository"
	"github.com/adityasuryadi/go-shop/services/auth/internal/usecase"
)

func main() {
	// viper := config.NewViper()
	logger := logger.NewLogger()

	_, cancel := context.WithCancel(context.Background())
	path := "./../../"

	configViper := config.NewViper(path)
	channel, _ := config.NewRabbitMqChannell(configViper, logger)

	jwtConfig := config.NewJWT(configViper)
	db := config.NewDatabase(configViper)
	validation := config.NewValidation(db)
	userRepository := repository.NewUserRespository(logger)
	redisClient := config.NewRedis(configViper)
	authUsecase := usecase.NewAuthUsecase(userRepository, db, jwtConfig, logger, validation, redisClient)
	userHandler := messaging.NewUserConsumer(logger, authUsecase)
	consumeOrderCfg := messaging.ConsumerConfig{
		Exchange:       "user.created",
		QueueName:      "user.create",
		RoutingKey:     "create",
		ConsumerTag:    "",
		BindingKey:     "create",
		WorkerPoolSize: 5,
	}

	go messaging.Consume(consumeOrderCfg, channel, userHandler.Consume)

	logger.Info("Starting worker...")

	terminateSignals := make(chan os.Signal, 1)
	signal.Notify(terminateSignals, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)

	stop := false
	for !stop {
		select {
		case s := <-terminateSignals:
			logger.Info("Got one of stop signals, shutting down worker gracefully, SIGNAL NAME :", s)
			cancel()
			stop = true
		}
	}

	time.Sleep(5 * time.Second)

}
