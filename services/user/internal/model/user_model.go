package model

type UserResponse struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type CreateUserRequest struct {
	Email                string `json:"email" validate:"required,email,unique=users"`
	FirstName            string `json:"first_name" validate:"required"`
	LastName             string `json:"last_name" validate:"required"`
	Phone                string `json:"phone" validate:"required"`
	Password             string `json:"password" validate:"required,min=4"`
	PasswordConfirmation string `json:"password_confirmation" validate:"required,min=4,eqfield=Password"`
}
