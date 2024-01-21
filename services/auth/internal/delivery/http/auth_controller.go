package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/adityasuryadi/go-shop/pkg/logger"
	middleware "github.com/adityasuryadi/go-shop/services/auth/internal/delivery"
	"github.com/adityasuryadi/go-shop/services/auth/internal/model"
	"github.com/adityasuryadi/go-shop/services/auth/internal/usecase"
	"github.com/go-chi/chi/v5"
)

type AuthController struct {
	Router  *chi.Mux
	Usecase usecase.AuthUsecase
	logger  *logger.Logger
}

func NewAuthController(r *chi.Mux, authUsecase usecase.AuthUsecase, log *logger.Logger) AuthController {
	return AuthController{
		Router:  r,
		Usecase: authUsecase,
		logger:  log,
	}
}

func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	payload := new(model.LoginRequest)
	if err := decoder.Decode(&payload); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Bad Request")
		return
	}
	accessToken, err := c.Usecase.Login(payload)
	if err != nil {
		fmt.Println("failed create token", err)
	}
	json.NewEncoder(w).Encode(model.WebResponse[*model.LoginResponse]{
		Code:   http.StatusOK,
		Status: "OK",
		Data: &model.LoginResponse{
			AccessToken: accessToken,
		},
	})
}

func (c *AuthController) Test(w http.ResponseWriter, r *http.Request) {
	headers := r.Header
	fmt.Println(headers)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"ok": "test"})
}

func (c *AuthController) Check(w http.ResponseWriter, r *http.Request) {
	headers := r.Header
	fmt.Println(headers)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"ok": "check"})
}

func (c *AuthController) InitRoute(Router *chi.Mux) {
	Router.Post("/login", c.Login)

	Router.Group(func(r chi.Router) {
		r.Use(middleware.Verify)
		r.Get("/check", c.Check)
		r.Get("/test", c.Test)
	})
}
