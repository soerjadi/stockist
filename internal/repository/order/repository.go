package order

import (
	"context"

	"github.com/soerjadi/stockist/internal/model"
	"github.com/soerjadi/stockist/internal/pkg/log"
	"github.com/soerjadi/stockist/internal/pkg/log/logger"
)

func (r orderRepository) CreateOrder(ctx context.Context, req model.Order) (model.Order, error) {
	var (
		orderModel model.Order
		err        error
	)

	if err = r.query.createOrder.GetContext(
		ctx,
		&orderModel,
		req.UserID,
		req.StoreID,
		req.TotalPrice,
		req.TotalAmount,
		req.Status,
	); err != nil {
		log.Errorw("[repository.order.CreateOrder] failed save order", logger.KV{
			"err": err,
			"req": req,
		})
		return model.Order{}, err
	}

	return orderModel, nil
}

func (r orderRepository) CreateOrderItem(ctx context.Context, req model.OrderItem) error {
	if _, err := r.query.createOrderItem.ExecContext(
		ctx,
		req.OrderID,
		req.ProductID,
		req.Amount,
		req.Price,
	); err != nil {
		log.Errorw("[repository.order.CreateOrderItem] failed save order item", logger.KV{
			"err":  err,
			"item": req,
		})
		return err
	}

	return nil
}

func (r orderRepository) UpdateOrderStatus(ctx context.Context, orderID int64, status string) error {
	if _, err := r.query.updateOrderStatus.ExecContext(
		ctx,
		status,
		orderID,
	); err != nil {
		log.Errorw("[repository.order.UpdateOrderStatus] failed update order item", logger.KV{
			"err": err,
		})

		return err
	}
	return nil
}
