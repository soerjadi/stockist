package order

import (
	"net/http"

	"github.com/soerjadi/stockist/internal/config"
	"github.com/soerjadi/stockist/internal/delivery/rest"
	"github.com/soerjadi/stockist/internal/delivery/rest/middleware"
	"github.com/soerjadi/stockist/internal/usecase/order"
	"github.com/soerjadi/stockist/internal/usecase/user"

	"github.com/gorilla/mux"
)

func NewHandler(usecase order.Usecase, userUsecase user.Usecase, config *config.Config) rest.API {
	return &Handler{
		usecase:     usecase,
		userUsecase: userUsecase,
		config:      config,
	}
}

func (h *Handler) RegisterRoutes(r *mux.Router) {
	orderHandler := r.PathPrefix("/order").Subrouter()
	orderHandler.Use(mux.CORSMethodMiddleware(r))
	orderHandler.Use(middleware.OnlyLoggedInUser(h.userUsecase, h.config))

	orderHandler.HandleFunc("/", rest.HandlerFunc(h.order).Serve).Methods(http.MethodPost)
}
