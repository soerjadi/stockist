package order

import (
	"github.com/redis/go-redis/v9"
	"github.com/soerjadi/stockist/internal/repository/order"
	"github.com/soerjadi/stockist/internal/repository/product"
)

func GetUsecase(repository order.Repository, productRepo product.Repository, redis *redis.Client) Usecase {
	return &orderUsecase{
		repository:  repository,
		productRepo: productRepo,
		redis:       redis,
	}
}
