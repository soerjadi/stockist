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
	Name                 string `json:"name" validate:"required"`
	Email                string `json:"email" validate:"required,email"`
	PhoneNumber          string `json:"phone_number" validate:"required"`
	Address              string `json:"address" validate:"required"`
	Password             string `json:"password" validate:"required"`
	PasswordConfirmation string `json:"password_confirmation" validate:"required,eqfield=Password"`
	Role                 string `json:"-"`
	Salt                 string `json:"-"`
}

type MapErrorRequest map[string]map[string]string

func (UserRequest) ErrorMessages(name string) map[string]string {
	return MapErrorRequest{
		"name": {
			"required": "Name field is required",
		},
		"email": {
			"required": "Email field is required",
			"email":    "Wrong email format",
		},
		"phone_number": {
			"required": "Phone number field is required",
		},
		"address": {
			"required": "Address field is required",
		},
		"password": {
			"required": "Password field is required",
		},
		"password_confirmation": {
			"required": "Password Confirmation is required",
			"eqfield":  "Password Confirmation doesn't equal with Password field",
		},
	}[name]
}

func (UserRequest) FieldName(name string) string {
	return map[string]string{
		"Name":                 "name",
		"Email":                "email",
		"PhoneNumber":          "phone_number",
		"Address":              "address",
		"Password":             "password",
		"PasswordConfirmation": "password_confirmation",
	}[name]
}

type UserLoginRequest struct {
	UserField string `json:"phone_or_email" validate:"required"`
	Password  string `json:"password" validate:"password"`
}
