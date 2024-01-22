package model

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"min=4"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}
