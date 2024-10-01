package product

import (
	"context"

	"github.com/soerjadi/stockist/internal/model"
	"github.com/soerjadi/stockist/internal/repository/product"
)

//go:generate mockgen -package=mocks -mock_names=Usecase=MockProductUsecase -destination=../../mocks/product_usecase_mock.go -source=type.go
type Usecase interface {
	GetByID(ctx context.Context, id int64) (model.Product, error)
	GetList(ctx context.Context, limit, offset int64) ([]model.Product, error)
	CreateProduct(ctx context.Context, req model.CreateProductRequest) (model.Product, error)
}

type productUsecase struct {
	repository product.Repository
}
