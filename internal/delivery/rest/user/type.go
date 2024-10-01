package user

import (
	"github.com/go-playground/validator/v10"
	"github.com/soerjadi/stockist/internal/usecase/user"
)

type Handler struct {
	usecase  user.Usecase
	validate *validator.Validate
}
