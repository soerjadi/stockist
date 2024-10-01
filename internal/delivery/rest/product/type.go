package product

import (
	"github.com/go-playground/validator/v10"
	"github.com/soerjadi/stockist/internal/config"
	"github.com/soerjadi/stockist/internal/usecase/product"
	"github.com/soerjadi/stockist/internal/usecase/store"
	"github.com/soerjadi/stockist/internal/usecase/user"
)

type Handler struct {
	usecase      product.Usecase
	userUsecase  user.Usecase
	storeUsecase store.Usecase
	validate     validator.Validate
	config       *config.Config
}
