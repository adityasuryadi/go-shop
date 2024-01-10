package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	handlerHttp "github.com/adityasuryadi/go-shop/services/user/internal/delivery/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	userController := handlerHttp.NewUserController(router)
	router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		// an example API handler
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})
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
