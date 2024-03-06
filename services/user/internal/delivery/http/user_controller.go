package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/adityasuryadi/go-shop/pkg/exception"
	"github.com/adityasuryadi/go-shop/pkg/logger"
	"github.com/adityasuryadi/go-shop/services/user/internal/config"
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
	validate := new(config.Validation)

	if err != nil && err.Status == exception.ERRRBADREQUEST {
		errResponse := model.ErrorResponse[any]{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
			Error:  validate.ErrorJson(err.Errors),
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errResponse)
		return
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
	vars := mux.Vars(r)
	userId := vars["user_id"]
	user, err := c.Usecase.FindById(userId)
	if err != nil {
		c.logger.Errorf("failed get user", err)
	}

	if user == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(model.WebResponse[*model.UserResponse]{
			Code:   http.StatusNotFound,
			Status: "NOT_FOUND",
			Data:   nil,
			Paging: nil,
		})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.WebResponse[*model.UserResponse]{
		Code:   200,
		Status: "OK",
		Data:   user,
		Paging: nil,
	})
	return
}

func (c *UserController) InitRoute(Router *mux.Router) {
	Router.HandleFunc("/user/{user_id}", c.Get).Methods("GET")
	Router.HandleFunc("/user", c.Create).Methods("POST")
}
