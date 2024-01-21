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
	db := config.NewDatabase(configViper)
	router := chi.NewRouter()
	userRepository := repository.NewUserRespository()
	authUsecase := usecase.NewAuthUsecase(userRepository, db, jwtConfig, logger)
	authController := handlerHttp.NewAuthController(router, authUsecase, logger)
	authController.InitRoute(router)

	// router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("data"))))
	// router.HandleFunc("/cek", func(w http.ResponseWriter, r *http.Request) {
	// 	dir, _ := os.Getwd()
	// 	fmt.Println("current path :" + dir)
	// 	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	// })
	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8001",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
