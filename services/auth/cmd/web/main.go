package main

import (
	"log"
	"net/http"
	"time"

	"github.com/adityasuryadi/go-shop/pkg/logger"
	"github.com/adityasuryadi/go-shop/services/auth/internal/config"
	handlerHttp "github.com/adityasuryadi/go-shop/services/auth/internal/delivery/http"
	"github.com/adityasuryadi/go-shop/services/auth/internal/repository"
	"github.com/adityasuryadi/go-shop/services/auth/internal/usecase"
	"github.com/go-chi/chi/v5"
)

func main() {
	configViper := config.NewViper()
	logger := logger.NewLogger()
	jwtConfig := config.NewJWT(configViper)
	validation := config.NewValidation()
	db := config.NewDatabase(configViper)
	router := chi.NewRouter()
	userRepository := repository.NewUserRespository()
	redisClient := config.NewRedis(configViper)
	authUsecase := usecase.NewAuthUsecase(userRepository, db, jwtConfig, logger, validation, redisClient)
	authController := handlerHttp.NewAuthController(router, authUsecase, logger)
	authController.InitRoute(router)

	fs := http.FileServer(http.Dir("data"))
	router.Handle("/static/*", http.StripPrefix("/static/", fs))
	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8001",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
