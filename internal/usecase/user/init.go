package user

import (
	"github.com/go-playground/validator/v10"
	"github.com/soerjadi/stockist/internal/repository/user"
)

func GetUsecase(repository user.Repository, validate *validator.Validate) Usecase {
	return &userUsecase{
		repository: repository,
		validate:   validate,
	}
}
