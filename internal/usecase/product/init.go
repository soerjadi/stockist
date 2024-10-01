package product

import "github.com/soerjadi/stockist/internal/repository/product"

func GetUsecase(repository product.Repository) Usecase {
	return &productUsecase{
		repository: repository,
	}
}
