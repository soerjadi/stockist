package user

import (
	"context"

	"github.com/soerjadi/stockist/internal/model"
	"github.com/soerjadi/stockist/internal/pkg/log"
	"github.com/soerjadi/stockist/internal/pkg/log/logger"
)

func (r userRepository) InsertUser(ctx context.Context, req model.UserRequest) (model.User, error) {
	var (
		user model.User
		err  error
	)

	if err = r.query.insertUser.GetContext(
		ctx,
		&user,
		req.Name,
		req.Email,
		req.PhoneNumber,
		req.Address,
		req.Role,
		req.Password,
		req.Salt,
	); err != nil {
		log.Errorw("[repository.user.InsertUser] failed save request", logger.KV{
			"err": err,
			"req": req,
		})

		return model.User{}, err
	}

	return user, nil

}

func (r userRepository) GetByID(ctx context.Context, id int64) (model.User, error) {
	var (
		userModel model.User
	)

	err := r.query.getUserByID.GetContext(ctx, &userModel, id)
	if err != nil {
		log.Errorw("[repository.user.getUserByID] got an error, %v", logger.KV{
			"err": err,
			"id":  id,
		})
		return model.User{}, err
	}

	return userModel, err
}

func (r userRepository) GetByPhoneNumber(ctx context.Context, phoneNumber string) (model.User, error) {
	var userModel model.User

	err := r.query.getUserByPhoneNumber.GetContext(ctx, &userModel, phoneNumber)
	if err != nil {
		log.Errorw("[repository.user.GetUserByPhoneNumber] got an error ", logger.KV{
			"err":         err,
			"phoneNumber": phoneNumber,
		})
		return model.User{}, err
	}

	return userModel, nil
}

func (r userRepository) GetByEmail(ctx context.Context, email string) (model.User, error) {
	var userModel model.User

	err := r.query.getUserByEmail.GetContext(ctx, &userModel, email)
	if err != nil {
		log.Errorw("[repository.user.GetUserByEmail] got an error ", logger.KV{
			"err":   err,
			"email": email,
		})
		return model.User{}, err
	}

	return userModel, nil
}
