package user

import (
	"context"

	"github.com/soerjadi/stockist/internal/config"
	"github.com/soerjadi/stockist/internal/model"
	"github.com/soerjadi/stockist/internal/repository/user"
)

//go:generate mockgen -package=mocks -mock_names=Usecase=MockUserUsecase -destination=../../mocks/user_usecase_mock.go -source=type.go
type Usecase interface {
	RegisterUser(ctx context.Context, req model.UserRequest) (model.User, error)
	GetByID(ctx context.Context, id int64) (model.User, error)
	Login(ctx context.Context, req model.UserLoginRequest) (string, string, error)
}

type userUsecase struct {
	repository user.Repository
	config     *config.Config
}
