package order

import (
	"encoding/json"
	"net/http"

	"github.com/soerjadi/stockist/internal/model"
	"github.com/soerjadi/stockist/internal/model/constant"
	"github.com/soerjadi/stockist/internal/pkg/log"
	"github.com/soerjadi/stockist/internal/pkg/log/logger"
)

func (h *Handler) order(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	dec := json.NewDecoder(r.Body)

	req := model.CreateOrderRequest{}

	if err := dec.Decode(&req); err != nil {
		log.Errorw("[delivery.rest.order.Order] fail decode", logger.KV{
			"err": err,
		})

		return nil, err
	}

	// if err := validator.Validate(r.Context(), &h.validate, req); err != nil {
	// 	log.Errorw("[delivery.rest.order.order] failed in validator", logger.KV{
	// 		"err": err,
	// 		"req": req,
	// 	})

	// 	return nil, err
	// }

	var currentUserID int64
	currentUserIDLog := r.Context().Value(constant.USER_ID_KEY_RESPONDENT)
	if currentUserIDLog != nil {
		currentUserID = currentUserIDLog.(int64)
	}

	req.UserID = currentUserID

	order, err := h.usecase.CreateOrder(r.Context(), req)
	if err != nil {
		return nil, err
	}

	return order, nil
}
