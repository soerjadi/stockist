package store

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/soerjadi/stockist/internal/model"
	"github.com/soerjadi/stockist/internal/pkg/log"
	"github.com/soerjadi/stockist/internal/pkg/log/logger"
	"github.com/soerjadi/stockist/internal/pkg/validator"
)

func (h Handler) register(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	dec := json.NewDecoder(r.Body)

	req := model.RegisterStoreRequest{}

	if err := dec.Decode(&req); err != nil {
		log.Errorw("[delivery.rest.store.register] faild decode request body", logger.KV{
			"err": err,
			"req": req,
		})

		return nil, errors.New("fail decode request body")
	}

	if err := validator.Validate(r.Context(), &h.validate, req); err != nil {
		log.Errorw("[delivery.rest.store.register] failed in validator", logger.KV{
			"err": err,
			"req": req,
		})

		return nil, err
	}

	store, err := h.usecase.CreateStore(r.Context(), req)
	if err != nil {
		log.Errorw("[delivery.rest.store.register] failed save store", logger.KV{
			"err": err,
			"req": req,
		})

		return nil, err
	}

	return store, nil
}

func (h Handler) detail(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return nil, err
	}

	store, err := h.usecase.GetByID(r.Context(), id)
	if err != nil {
		log.Errorw("[delivery.rest.store.detail] failed get by id", logger.KV{
			"err": err,
			"id":  id,
		})
		return nil, err
	}

	return store, nil
}
