package store

import (
	"context"

	"github.com/soerjadi/stockist/internal/model"
	"github.com/soerjadi/stockist/internal/pkg/log"
	"github.com/soerjadi/stockist/internal/pkg/log/logger"
)

func (r storeRepository) InsertStore(ctx context.Context, req model.RegisterStoreRequest) (model.Store, error) {
	var (
		store model.Store
		err   error
	)

	if err = r.query.insertStore.GetContext(
		ctx,
		&store,
		req.Name,
		req.Description,
		req.Address,
		req.OwnerID,
	); err != nil {
		log.Errorw("[repository.store.InsertStore] failed save store", logger.KV{
			"err": err,
			"req": req,
		})
		return model.Store{}, err
	}

	return store, nil
}

func (r storeRepository) GetByID(ctx context.Context, id int64) (model.Store, error) {
	var storeModel model.Store

	err := r.query.getByID.GetContext(ctx, &storeModel, id)
	if err != nil {
		log.Errorw("[repository.store.GetByID] failed get by id", logger.KV{
			"err": err,
		})
		return model.Store{}, err
	}

	return storeModel, nil
}
