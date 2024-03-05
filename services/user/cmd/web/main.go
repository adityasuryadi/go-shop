package main

import (
	"log"
	"net/http"
	"time"

	"github.com/adityasuryadi/go-shop/pkg/logger"
	"github.com/adityasuryadi/go-shop/services/user/internal/config"
	handlerHttp "github.com/adityasuryadi/go-shop/services/user/internal/delivery/http"
	"github.com/adityasuryadi/go-shop/services/user/internal/repository"
	"github.com/adityasuryadi/go-shop/services/user/internal/usecase"

	"github.com/gorilla/mux"
)

func main() {
	configViper := config.NewViper()
	logger := logger.NewLogger()
	channel, err := config.NewRabbitMqChannell(configViper, logger)
	if err != nil {
		logger.Error("failed to create channel", err)
	}
	db := config.NewDatabase(configViper)
	router := mux.NewRouter()
	userRepository := repository.NewUserRepository(logger)
	userUsecase := usecase.NewUserUsecase(db, userRepository, channel)
	userController := handlerHttp.NewUserController(router, userUsecase, logger)
	defer channel.Close()
	userController.InitRoute(router)
	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
