package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/adityasuryadi/go-shop/pkg/exception"
	"github.com/adityasuryadi/go-shop/pkg/logger"
	"github.com/adityasuryadi/go-shop/services/auth/internal/config"
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

func (c *AuthController) Register(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	request := new(model.RegisterRequest)

	if err := decoder.Decode(request); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Bad Request")
		return
	}

	validation := new(config.Validation)
	err := c.Usecase.Register(request)

	if err != nil && err.Status == exception.ERRRBADREQUEST {
		errValidation := validation.ErrorJson(err.Errors)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse[any]{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
			Error:  errValidation,
		})
		return
	}

	if err != nil && err.Status == exception.ERRBUSSINESS {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse[any]{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL_SERVER_ERROR",
			Error:  err.Errors,
		})
		return
	}

	json.NewEncoder(w).Encode(model.WebResponse[any]{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	})
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

	validation := new(config.Validation)
	response, err := c.Usecase.Login(payload)
	if err != nil && err.Status == exception.ERRRBADREQUEST {
		errValidation := validation.ErrorJson(err.Errors)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse[any]{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
			Error:  errValidation,
		})
		return
	}

	if err != nil && err.Status == exception.ERRNOTFOUND {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		fmt.Println("error ", err.Errors)
		json.NewEncoder(w).Encode(model.ErrorResponse[any]{
			Code:   http.StatusNotFound,
			Status: "NOT_FOUND",
			Error:  err.Errors.Error(),
		})
		return
	}
	cookie := &http.Cookie{}
	cookie.HttpOnly = true
	cookie.Name = "refresh_token"
	cookie.Value = response.RefreshToken
	cookie.Expires = time.Now().Add(7200 * time.Hour)
	http.SetCookie(w, cookie)
	json.NewEncoder(w).Encode(model.WebResponse[*model.LoginResponse]{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (c *AuthController) RefreshToken(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("refresh_token")
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse[any]{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
			Error:  "UNAUTHORIZED",
		})
		return
	}

	response, customErr := c.Usecase.RefreshToken(cookie.Value)
	if customErr != nil {
		json.NewEncoder(w).Encode(model.ErrorResponse[any]{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
			Error:  "UNAUTHORIZED",
		})
		return
	}

	json.NewEncoder(w).Encode(model.WebResponse[*model.LoginResponse]{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   response,
	})

}

func (c *AuthController) Test(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{}
	cookie.HttpOnly = true
	cookie.Value = "test cookie"
	cookie.Expires = time.Now().Add(7200 * time.Hour)

	headers := r.Header
	fmt.Println(headers)

	http.SetCookie(w, cookie)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"ok": "test"})
}

func (c *AuthController) Check(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"ok": "check"})
}

func (c *AuthController) Logout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jwtToken := r.Header.Get("Authorization")
	err := c.Usecase.Logout(jwtToken)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(model.ErrorResponse[any]{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
			Error:  "UNAUTHORIZED",
		})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.WebResponse[any]{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	})
}

func (c *AuthController) InitRoute(Router *chi.Mux) {
	Router.Post("/register", c.Register)
	Router.Post("/login", c.Login)
	Router.Post("/refreshtoken", c.RefreshToken)
	Router.Group(func(r chi.Router) {
		r.Use(middleware.Verify)
		r.Get("/check", c.Check)
		r.Get("/test", c.Test)
		r.Post("/logout", c.Logout)
	})
}
