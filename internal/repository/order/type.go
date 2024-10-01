package order

import (
	"context"

	"github.com/soerjadi/stockist/internal/model"
)

//go:generate mockgen -package=mocks -mock_names=Repository=MockOrderRepository -destination=../../mocks/order_repo_mock.go -source=type.go
type Repository interface {
	CreateOrder(ctx context.Context, req model.Order) (model.Order, error)
	CreateOrderItem(ctx context.Context, req model.OrderItem) error
	UpdateOrderStatus(ctx context.Context, orderID int64, status string) error
}

type orderRepository struct {
	query prepareQuery
}
