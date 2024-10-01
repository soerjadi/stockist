package product

import (
	"context"

	"github.com/soerjadi/stockist/internal/model"
	"github.com/soerjadi/stockist/internal/pkg/log"
	"github.com/soerjadi/stockist/internal/pkg/log/logger"
)

func (u productUsecase) GetByID(ctx context.Context, id int64) (model.Product, error) {
	product, err := u.repository.GetByID(ctx, id)
	if err != nil {
		log.Errorw("[usecase.product.GetByID] failed get by id", logger.KV{
			"err": err,
		})

		return model.Product{}, err
	}

	return product, nil
}

func (u productUsecase) GetList(ctx context.Context, limit, offset int64) ([]model.Product, error) {
	products, err := u.repository.GetList(ctx, offset, limit)
	if err != nil {
		log.Errorw("[usecase.product.GetList] failed get list by id", logger.KV{
			"err": err,
		})

		return nil, err
	}

	return products, nil
}

func (u productUsecase) CreateProduct(ctx context.Context, req model.CreateProductRequest) (model.Product, error) {
	product, err := u.repository.CreateProduct(ctx, req)
	if err != nil {
		log.Errorw("[usecase.product.CreateProduct] fail create product", logger.KV{
			"err": err,
			"req": req,
		})
		return model.Product{}, err
	}

	return product, nil
}
