package store

import (
	"github.com/go-playground/validator/v10"
	"github.com/soerjadi/stockist/internal/config"
	"github.com/soerjadi/stockist/internal/usecase/store"
	"github.com/soerjadi/stockist/internal/usecase/user"
)

type Handler struct {
	usecase     store.Usecase
	userUsecase user.Usecase
	validate    validator.Validate
	config      *config.Config
}
