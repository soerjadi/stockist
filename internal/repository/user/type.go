package user

import (
	"context"

	"github.com/soerjadi/stockist/internal/model"
)

//go:generate mockgen -package=mocks -mock_names=Repository=MockUserRepository -destination=../../mocks/user_repo_mock.go -source=type.go
type Repository interface {
	InsertUser(ctx context.Context, req model.UserRequest) (model.User, error)
	GetByID(ctx context.Context, id int64) (model.User, error)
	GetByPhoneNumber(ctx context.Context, phoneNumber string) (model.User, error)
	GetByEmail(ctx context.Context, email string) (model.User, error)
}

type userRepository struct {
	query prepareQuery
}
