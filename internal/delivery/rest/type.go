package rest

import "github.com/gorilla/mux"

type API interface {
	RegisterRoutes(router *mux.Router)
}

type Response struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message,omitempty"`
	Status  string      `json:"status"`
}
