package product

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/soerjadi/stockist/internal/model"
	"github.com/soerjadi/stockist/internal/model/constant"
	"github.com/soerjadi/stockist/internal/pkg/log"
	"github.com/soerjadi/stockist/internal/pkg/log/logger"
	"github.com/soerjadi/stockist/internal/pkg/validator"
)

func (h *Handler) getByID(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return nil, err
	}

	product, err := h.usecase.GetByID(r.Context(), id)
	if err != nil {
		log.Errorw("[delivery.rest.product.getByID] failed get by id", logger.KV{
			"err": err,
			"id":  id,
		})
		return nil, err
	}

	return product, nil
}

func (h Handler) getList(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	offsetStr := r.URL.Query().Get("offset")
	if offsetStr == "" {
		offsetStr = "0"
	}
	limitStr := r.URL.Query().Get("limit")
	if limitStr == "" {
		limitStr = "10"
	}

	offset, err := strconv.ParseInt(offsetStr, 10, 64)
	if err != nil {
		return nil, err
	}
	limit, err := strconv.ParseInt(limitStr, 10, 64)
	if err != nil {
		return nil, err
	}

	products, err := h.usecase.GetList(r.Context(), limit, offset)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (h Handler) CreateProduct(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	dec := json.NewDecoder(r.Body)

	req := model.CreateProductRequest{}

	if err := dec.Decode(&req); err != nil {
		log.Errorw("[delivery.rest.product.CreateProduct] failed decode request body", logger.KV{
			"err": err,
		})

		return nil, err
	}

	if err := validator.Validate(r.Context(), &h.validate, req); err != nil {
		log.Errorw("[delivery.rest.product.CreateProduct] failed in validator", logger.KV{
			"err": err,
			"req": req,
		})

		return nil, err
	}

	if err := h.validateProductStore(r.Context(), req); err != nil {
		return nil, err
	}

	product, err := h.usecase.CreateProduct(r.Context(), req)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (h Handler) validateProductStore(ctx context.Context, req model.CreateProductRequest) error {
	var currentUserID int64
	currentUserIDLog := ctx.Value(constant.USER_ID_KEY_RESPONDENT)
	if currentUserIDLog != nil {
		currentUserID = currentUserIDLog.(int64)
	}

	store, err := h.storeUsecase.GetByID(ctx, req.StoreID)
	if err != nil {
		log.Errorw("[delivery.rest.product.validateProductStore] fail validate store", logger.KV{
			"err": err,
		})
		return err
	}

	if store.ManagerID != currentUserID {
		return errors.New("only owner store able to create product")
	}

	return nil
}
