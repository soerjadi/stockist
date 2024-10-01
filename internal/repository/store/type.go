package store

import (
	"context"

	"github.com/soerjadi/stockist/internal/model"
)

//go:generate mockgen -package=mocks -mock_names=Repository=MockStoreRepository -destination=../../mocks/store_repo_mock.go -source=type.go
type Repository interface {
	InsertStore(ctx context.Context, req model.RegisterStoreRequest) (model.Store, error)
	GetByID(ctx context.Context, id int64) (model.Store, error)
}

type storeRepository struct {
	query prepareQuery
}
