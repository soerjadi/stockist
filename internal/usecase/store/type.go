package store

import (
	"context"

	"github.com/soerjadi/stockist/internal/model"
	"github.com/soerjadi/stockist/internal/repository/store"
)

//go:generate mockgen -package=mocks -mock_names=Usecase=MockStoreUsecase -destination=../../mocks/store_usecase_mock.go -source=type.go
type Usecase interface {
	CreateStore(ctx context.Context, req model.RegisterStoreRequest) (model.Store, error)
	GetByID(ctx context.Context, id int64) (model.Store, error)
}

type storeUsecase struct {
	repository store.Repository
}
