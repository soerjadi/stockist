package user

import (
	"github.com/soerjadi/stockist/internal/config"
	"github.com/soerjadi/stockist/internal/repository/user"
)

func GetUsecase(repository user.Repository, config *config.Config) Usecase {
	return &userUsecase{
		repository: repository,
		config:     config,
	}
}
