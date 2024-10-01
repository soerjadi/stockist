package order

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/soerjadi/stockist/internal/model"
	"github.com/soerjadi/stockist/internal/pkg/log"
	"github.com/soerjadi/stockist/internal/pkg/log/logger"
)

func (u orderUsecase) CreateOrder(ctx context.Context, req model.CreateOrderRequest) (model.Order, error) {
	var (
		orderModel model.Order
	)

	orderModel.UserID = req.UserID
	orderModel.StoreID = req.StoreID
	orderModel.Status = model.ORDER_STATUS_CREATED

	// check stock from product item
	for _, _product := range req.Products {
		product, err := u.productRepo.GetByID(ctx, _product.ProductID)
		if err != nil {
			log.Errorw("[usecase.order.CreateOrder] product with id doesn't exists", logger.KV{})
			return model.Order{}, fmt.Errorf("product with id %d doesn't exists", _product.ProductID)
		}

		// check is there any reserved stock
		if err = u.reserveStock(ctx, product, req.UserID, _product.Total); err != nil {
			return model.Order{}, err
		}
	}

	// placing an order
	order, err := u.repository.CreateOrder(ctx, orderModel)
	if err != nil {
		return model.Order{}, err
	}

	go func() {
		for _, product := range req.Products {
			item := model.OrderItem{
				OrderID:   order.ID,
				ProductID: product.ProductID,
				Amount:    product.Total,
				Price:     product.Price,
			}
			err := u.repository.CreateOrderItem(ctx, item)
			if err != nil {
				log.Errorw("[usecase.order.CreateOrder] failed save order item", logger.KV{
					"err":     err,
					"product": product,
					"order":   order,
				})
				return
			}
		}
	}()

	return order, nil
}

func (u orderUsecase) reserveStock(ctx context.Context, product model.Product, userID, totalReserve int64) error {
	// check is there any reserve stock
	// key is user-product
	key := fmt.Sprintf("product-id-%d", product.ID)
	userKey := fmt.Sprintf("product-%d-user-%d", product.ID, userID)
	rsvStock, err := u.redis.Get(ctx, key).Result()
	if err != redis.Nil {
		log.Errorw("[usecase.order.reserveStock] get reserve product stock error", logger.KV{
			"err": err,
			"key": key,
		})

		rsvStock = "0"
	}

	rsvStockByUser, err := u.redis.Get(ctx, userKey).Result()
	if err != redis.Nil {
		log.Errorw("[usecase.order.reserveStock] get reserve stock by user error", logger.KV{
			"err": err,
			"key": key,
		})

		rsvStockByUser = "0"

	}

	// getting reserved stock filter by key product-id
	rsvStockInt, err := strconv.ParseInt(rsvStock, 10, 64)
	if err != nil {
		rsvStockInt = 0
	}

	// getting reserved stock filter by key product-id-user-id
	rsvStockByUserInt, err := strconv.ParseInt(rsvStockByUser, 10, 64)
	if err != nil {
		rsvStockByUserInt = 0
	}

	if product.Stock < (rsvStockInt + totalReserve) {
		log.Errorw("[usecase.order.reserveStock] stock is not enough", logger.KV{
			"product_id": product.ID,
			"stock":      product.Stock,
			"reserved":   rsvStockInt,
		})

		return fmt.Errorf("stock is not enough for product id %d", product.ID)
	}

	ttl := time.Duration(10) * time.Second
	u.redis.SetNX(ctx, rsvStock, (rsvStockInt + totalReserve), ttl)
	u.redis.SetNX(ctx, rsvStockByUser, (rsvStockByUserInt + totalReserve), ttl)

	return nil

}

func (u orderUsecase) Checkout(ctx context.Context, orderID int64) error {
	return nil
}

func (u orderUsecase) SetupPaymentMethod(ctx context.Context, orderID int64) error {
	return nil
}
