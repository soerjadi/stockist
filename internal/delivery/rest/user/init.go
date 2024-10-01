package user

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/soerjadi/stockist/internal/delivery/rest"
	"github.com/soerjadi/stockist/internal/usecase/user"

	"github.com/gorilla/mux"
)

func NewHandler(usecase user.Usecase, validate *validator.Validate) rest.API {
	return &Handler{
		usecase:  usecase,
		validate: validate,
	}
}

func (h *Handler) RegisterRoutes(r *mux.Router) {
	r.Use(mux.CORSMethodMiddleware(r))
	r.HandleFunc("/register", rest.HandlerFunc(h.register).Serve).Methods(http.MethodPost)
	r.HandleFunc("/login", rest.HandlerFunc(h.login).Serve).Methods(http.MethodPost)
}
