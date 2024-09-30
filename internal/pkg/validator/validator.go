package validator

import (
	"context"
	"errors"

	"github.com/go-playground/validator/v10"
)

type request interface {
	ErrorMessages(name string) map[string]string
	FieldName(name string) string
}

func Validate(ctx context.Context, validate *validator.Validate, req request) (err error) {
	if err = validate.StructCtx(ctx, req); err != nil {
		validationErr, ok := err.(validator.ValidationErrors)
		if !ok {
			return
		}

		if len(validationErr) > 0 {
			fieldError := validationErr[0]
			fieldName := req.FieldName(fieldError.Field())
			errMsg := req.ErrorMessages(fieldName)[fieldError.ActualTag()]

			err = errors.New(errMsg)

			return
		}

	}

	return
}
