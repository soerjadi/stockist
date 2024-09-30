package model

import "time"

type User struct {
	ID          int64     `db:"id"`
	Name        string    `db:"name"`
	Email       string    `db:"email"`
	PhoneNumber string    `db:"phone_number"`
	Address     string    `db:"address"`
	Role        string    `db:"role"`
	Password    string    `db:"password"`
	Salt        string    `db:"salt"`
	CreatedAt   time.Time `db:"created_at"`
}

const (
	USER_ROLE_ADMIN   = "admin"
	USER_ROLE_USER    = "user"
	USER_ROLE_SHOPPER = "shopper"
)
