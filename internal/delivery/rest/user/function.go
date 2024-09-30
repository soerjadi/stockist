package user

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/soerjadi/stockist/internal/model"
	"github.com/soerjadi/stockist/internal/pkg/log"
	"github.com/soerjadi/stockist/internal/pkg/log/logger"
)

func (h Handler) register(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	dec := json.NewDecoder(r.Body)

	req := model.UserRequest{}

	if err := dec.Decode(&req); err != nil {
		log.Errorw("[delivery.rest.user.register] failed decode request body", logger.KV{
			"err": err,
		})
		return nil, errors.New("fail to decode request body")
	}

	user, err := h.usecase.RegisterUser(r.Context(), req)
	if err != nil {
		log.Errorw("[delivery.rest.user.register] failed save user", logger.KV{
			"err": err,
		})
		return nil, err
	}

	return user, nil
}

func (h Handler) login(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	dec := json.NewDecoder(r.Body)

	req := model.UserLoginRequest{}

	if err := dec.Decode(&req); err != nil {
		log.Errorw("[delivery.rest.user.login] failed decode request body", logger.KV{
			"err": err,
		})
		return nil, errors.New("fail decode request body")
	}

	token, refresh, err := h.usecase.Login(r.Context(), req)
	if err != nil {
		log.Errorw("[deliver.rest.user.login] failed login", logger.KV{
			"err": err,
			"req": req,
		})

		return nil, err
	}

	return map[string]string{
		"access_token":  token,
		"refresh_token": refresh,
	}, nil
}
