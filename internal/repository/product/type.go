package product

import (
	"context"

	"github.com/soerjadi/stockist/internal/model"
)

//go:generate mockgen -package=mocks -mock_names=Repository=MockProductRepository -destination=../../mocks/product_repo_mock.go -source=type.go
type Repository interface {
	GetByID(ctx context.Context, id int64) (model.Product, error)
	GetList(ctx context.Context, offset, limit int64) ([]model.Product, error)
	CreateProduct(ctx context.Context, req model.CreateProductRequest) (model.Product, error)
}

type productRepository struct {
	query prepareQuery
}
