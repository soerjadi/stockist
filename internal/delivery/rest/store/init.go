package store

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/soerjadi/stockist/internal/config"
	"github.com/soerjadi/stockist/internal/delivery/rest"
	"github.com/soerjadi/stockist/internal/delivery/rest/middleware"
	"github.com/soerjadi/stockist/internal/usecase/store"
	"github.com/soerjadi/stockist/internal/usecase/user"

	"github.com/gorilla/mux"
)

func NewHandler(usecase store.Usecase, userUsecase user.Usecase, validate *validator.Validate, cfg *config.Config) rest.API {
	return &Handler{
		usecase:     usecase,
		userUsecase: userUsecase,
		validate:    *validate,
		config:      cfg,
	}
}

func (h *Handler) RegisterRoutes(r *mux.Router) {
	storeHandler := r.PathPrefix("/store").Subrouter()
	storeHandler.Use(mux.CORSMethodMiddleware(r))
	storeHandler.Use(middleware.OnlyLoggedInUser(h.userUsecase, h.config))

	storeHandler.HandleFunc("/register", rest.HandlerFunc(h.register).Serve).Methods(http.MethodPost)
	// storeHandler.HandleFunc("/warehouse", rest.HandlerFunc(h.warehouse).Serve).Methods(http.MethodPost)
	storeHandler.HandleFunc("/{id:[0-9]+}", rest.HandlerFunc(h.detail).Serve).Methods(http.MethodGet)

}
