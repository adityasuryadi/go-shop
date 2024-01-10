package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/adityasuryadi/go-shop/pkg/logger"
	"github.com/adityasuryadi/go-shop/services/user/internal/model"
	"github.com/adityasuryadi/go-shop/services/user/internal/usecase"

	"github.com/gorilla/mux"
)

type UserController struct {
	Router  *mux.Router
	Usecase usecase.UserUsecase
	logger  *logger.Logger
}

func NewUserController(r *mux.Router, userUsecase usecase.UserUsecase, log *logger.Logger) UserController {
	return UserController{
		Router:  r,
		Usecase: userUsecase,
		logger:  log,
	}
}

func (c *UserController) Create(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	payload := new(model.CreateUserRequest)
	if err := decoder.Decode(&payload); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Bad Request")
		return
	}

	response, err := c.Usecase.Insert(context.Background(), payload)
	if err != nil {
		c.logger.Errorf("failed to create user ", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.WebResponse[*model.UserResponse]{
		Code:   200,
		Status: "OK",
		Data:   response,
		Paging: nil,
	})
}

func (c *UserController) Get(w http.ResponseWriter, r *http.Request) {
	var user model.UserResponse
	user = model.UserResponse{
		FirstName: "Aditya",
		LastName:  "Suryadi",
		Email:     "adit@mail.com",
		Phone:     "09372487234",
		CreatedAt: 123124,
		UpdatedAt: 124124,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.WebResponse[model.UserResponse]{
		Code:   200,
		Status: "OK",
		Data:   user,
		Paging: nil,
	})
}

func (c *UserController) InitRoute(Router *mux.Router) {
	Router.HandleFunc("/", c.Get).Methods("GET")
	Router.HandleFunc("/", c.Create).Methods("POST")
}
