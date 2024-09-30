package model

import "time"

type User struct {
	ID          int64     `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
	Email       string    `db:"email" json:"email"`
	PhoneNumber string    `db:"phone_number" json:"phone_number"`
	Address     string    `db:"address" json:"address"`
	Role        string    `db:"role" json:"-"`
	Password    string    `db:"password" json:"-"`
	Salt        string    `db:"salt" json:"-"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
}

const (
	USER_ROLE_ADMIN   = "admin"
	USER_ROLE_USER    = "user"
	USER_ROLE_SHOPPER = "shopper"
)

type UserRequest struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
	Password    string `json:"password"`
	Role        string `json:"-"`
	Salt        string `json:"-"`
}

type UserLoginRequest struct {
	UserField string `json:"phone_or_email" validate:"required"`
	Password  string `json:"password" validate:"password"`
}

type UserRegisterRequest struct {
	Name                 string `json:"name" validate:"required"`
	Email                string `json:"email" validate:"required,email"`
	PhoneNumber          string `json:"phone_number" validate:"required"`
	Address              string `json:"address" validate:"required"`
	Password             string `json:"password" validate:"required"`
	PasswordConfirmation string `json:"password_confirmation" validate:"required,eqfield=Password"`
}
