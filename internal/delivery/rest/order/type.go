package order

import (
	"github.com/go-playground/validator/v10"
	"github.com/soerjadi/stockist/internal/config"
	"github.com/soerjadi/stockist/internal/usecase/order"
	"github.com/soerjadi/stockist/internal/usecase/user"
)

type Handler struct {
	usecase     order.Usecase
	userUsecase user.Usecase
	config      *config.Config
	validate    validator.Validate
}
