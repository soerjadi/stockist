package store

import "github.com/soerjadi/stockist/internal/repository/store"

func GetUsecase(repository store.Repository) Usecase {
	return &storeUsecase{
		repository: repository,
	}
}
