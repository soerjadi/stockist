package product

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/soerjadi/stockist/internal/config"
	"github.com/soerjadi/stockist/internal/delivery/rest"
	"github.com/soerjadi/stockist/internal/delivery/rest/middleware"
	"github.com/soerjadi/stockist/internal/usecase/product"
	"github.com/soerjadi/stockist/internal/usecase/store"
	"github.com/soerjadi/stockist/internal/usecase/user"

	"github.com/gorilla/mux"
)

func NewHandler(
	usecase product.Usecase,
	userUsecase user.Usecase,
	storeUsecase store.Usecase,
	validate *validator.Validate,
	cfg *config.Config,
) rest.API {
	return &Handler{
		usecase:      usecase,
		userUsecase:  userUsecase,
		storeUsecase: storeUsecase,
		validate:     *validate,
		config:       cfg,
	}
}

func (h *Handler) RegisterRoutes(r *mux.Router) {
	productHandler := r.PathPrefix("/product").Subrouter()
	productHandler.Use(mux.CORSMethodMiddleware(r))
	productHandler.Use(middleware.OnlyLoggedInUser(h.userUsecase, h.config))

	productHandler.HandleFunc("/{id:[0-9]+}", rest.HandlerFunc(h.getByID).Serve).Methods(http.MethodGet)
	productHandler.HandleFunc("/list", rest.HandlerFunc(h.getList).Serve).Methods(http.MethodGet)
	productHandler.HandleFunc("", rest.HandlerFunc(h.CreateProduct).Serve).Methods(http.MethodPost)

	// productStoreHandler := r.PathPrefix("/store/{store:[0-9]+}").Subrouter()
	// productStoreHandler.Use(mux.CORSMethodMiddleware(r))
	// productStoreHandler.Use(middleware.OnlyLoggedInUser(h.userUsecase, h.config))

}
