package order

import (
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/soerjadi/stockist/internal/model"
	"github.com/soerjadi/stockist/internal/repository/order"
	"github.com/soerjadi/stockist/internal/repository/product"
)

//go:generate mockgen -package=mocks -mock_names=Usecase=MockOrderUsecase -destination=../../mocks/order_usecase_mock.go -source=type.go
type Usecase interface {
	CreateOrder(ctx context.Context, req model.CreateOrderRequest) (model.Order, error)
	Checkout(ctx context.Context, orderID int64) error
	SetupPaymentMethod(ctx context.Context, orderID int64) error
}

type orderUsecase struct {
	repository  order.Repository
	productRepo product.Repository
	redis       *redis.Client
}
