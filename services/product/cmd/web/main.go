package main

import (
	"log"
	"net/http"
	"time"

	"github.com/adityasuryadi/go-shop/pkg/logger"
	"github.com/adityasuryadi/go-shop/services/product/internal/config"
	handlerHttp "github.com/adityasuryadi/go-shop/services/product/internal/delivery/http"
	"github.com/adityasuryadi/go-shop/services/product/internal/repository"
	"github.com/adityasuryadi/go-shop/services/product/internal/usecase"
	"github.com/go-chi/chi/v5"
)

func main() {
	configViper := config.NewViper()
	router := chi.NewRouter()
	validation := config.NewValidation()
	config.NewDatabase(configViper)
	db := config.NewDatabase(configViper)
	productRepository := repository.NewProductRepository(db)
	productUsecase := usecase.NewProductUsecase(db, productRepository, validation)
	ProductController := handlerHttp.NewProductController(router, productUsecase, logger.NewLogger())
	ProductController.InitRoute(router)
	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8002",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Product Service"))
	})

	log.Fatal(srv.ListenAndServe())
}
