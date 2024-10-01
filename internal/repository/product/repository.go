package product

import (
	"context"

	"github.com/soerjadi/stockist/internal/model"
	"github.com/soerjadi/stockist/internal/pkg/log"
	"github.com/soerjadi/stockist/internal/pkg/log/logger"
)

func (r productRepository) GetByID(ctx context.Context, id int64) (model.Product, error) {
	var productModel model.Product

	err := r.query.getByID.GetContext(ctx, &productModel, id)
	if err != nil {
		log.Errorw("[repository.product.GetByID] failed product get by id", logger.KV{
			"err": err,
		})
		return model.Product{}, err
	}

	return productModel, nil
}

func (r productRepository) GetList(ctx context.Context, offset, limit int64) ([]model.Product, error) {
	var products []model.Product

	err := r.query.getList.SelectContext(ctx, &products, limit, offset)
	if err != nil {
		log.Errorw("[repository.product.GetList] failed get product list", logger.KV{
			"err": err,
		})
		return nil, err
	}

	return products, nil
}

func (r productRepository) CreateProduct(ctx context.Context, req model.CreateProductRequest) (model.Product, error) {
	var (
		product model.Product
		err     error
	)

	if err = r.query.createProduct.GetContext(
		ctx,
		&product,
		req.Name,
		req.Description,
		req.Weight,
		req.Price,
		req.StoreID,
		req.Stock,
		req.Images,
	); err != nil {
		log.Errorw("[repository.product.CreateProduc] failed store product", logger.KV{
			"err": err,
			"req": req,
		})

		return model.Product{}, err
	}

	return product, nil
}
