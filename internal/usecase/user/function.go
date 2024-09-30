package user

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/soerjadi/stockist/internal/model"
	"github.com/soerjadi/stockist/internal/model/constant"
	"github.com/soerjadi/stockist/internal/pkg/log"
	"github.com/soerjadi/stockist/internal/pkg/log/logger"
	"github.com/soerjadi/stockist/internal/pkg/str"
	"github.com/soerjadi/stockist/internal/pkg/token"
)

func (u userUsecase) RegisterUser(ctx context.Context, req model.UserRequest) (model.User, error) {
	salt := str.GenerateSalt()
	password := fmt.Sprintf("%s<<>>%s", req.Password, salt)

	passwordHash, err := str.HashStr(password)
	if err != nil {
		log.Errorw("[usecase.user.RegisterUser] failed generate hash password", logger.KV{
			"err": err,
		})
		return model.User{}, err
	}

	req.Password = passwordHash
	req.Salt = salt
	req.Role = model.USER_ROLE_USER

	result, err := u.repository.InsertUser(ctx, req)
	if err != nil {
		log.Errorw("[usecase.user.RegisterUser] failed insert user", logger.KV{
			"err":     err,
			"request": req,
		})
		return model.User{}, err
	}

	return result, nil
}

func (u userUsecase) GetByID(ctx context.Context, id int64) (model.User, error) {
	user, err := u.repository.GetByID(ctx, id)
	if err != nil {
		log.Errorw("[usecase.user.GetByID] failed get user by id", logger.KV{
			"err": err,
			"id":  id,
		})
		return model.User{}, err
	}

	return user, nil
}

func (u userUsecase) Login(ctx context.Context, req model.UserLoginRequest) (string, string, error) {
	var (
		user model.User
		err  error
	)
	userIdentifier := "phone_number"

	if len(strings.Split(req.UserField, "@")) > 1 {
		userIdentifier = "email"
	}

	switch userIdentifier {
	case "email":
		user, err = u.repository.GetByEmail(ctx, req.UserField)
	default:
		user, err = u.repository.GetByPhoneNumber(ctx, req.UserField)
	}

	if err != nil {
		log.Errorw("[usecase.user.Login] fail get user identity", logger.KV{
			"err":        err,
			"identifier": userIdentifier,
			"request":    req,
		})

		return "", "", errors.New(constant.ERROR_USER_NOT_FOUND_EN)
	}

	plainPassword := fmt.Sprintf("%s<<>>%s", req.Password, user.Salt)
	if !str.CompareHash(user.Password, plainPassword) {
		return "", "", errors.New(constant.ERROR_PASSWORD_NOT_MATCH_EN)
	}

	accessToken, err := token.GenerateAccessToken(user.ID)
	if err != nil {
		log.Errorw("[usecase.user.Login] failed generate access token", logger.KV{
			"err": err,
		})
		return "", "", err
	}

	refreshToken, err := token.GenerateRefreshToken(user.ID)
	if err != nil {
		log.Errorw("[usecase.user.Login] failed generate refresh token", logger.KV{
			"err": err,
		})
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
