package store

import (
	"context"

	"github.com/soerjadi/stockist/internal/model"
	"github.com/soerjadi/stockist/internal/model/constant"
	"github.com/soerjadi/stockist/internal/pkg/log"
	"github.com/soerjadi/stockist/internal/pkg/log/logger"
)

func (u storeUsecase) CreateStore(ctx context.Context, req model.RegisterStoreRequest) (model.Store, error) {
	var currentUserID int64
	currentUserIDLog := ctx.Value(constant.USER_ID_KEY_RESPONDENT)
	if currentUserIDLog != nil {
		currentUserID = currentUserIDLog.(int64)
	}

	req.OwnerID = currentUserID
	store, err := u.repository.InsertStore(ctx, req)
	if err != nil {
		log.Errorw("[usecase.store.CreateStore] fail create store", logger.KV{
			"err": err,
			"req": req,
		})
		return model.Store{}, err
	}

	return store, nil
}

func (u storeUsecase) GetByID(ctx context.Context, id int64) (model.Store, error) {
	store, err := u.repository.GetByID(ctx, id)
	if err != nil {
		log.Errorw("[usecase.store.GetByID] failed get store by id", logger.KV{
			"err": err,
			"id":  id,
		})

		return model.Store{}, err
	}

	return store, nil
}
