package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (h *HTTP) setRoutes(r *mux.Router) {
	api := r.PathPrefix("/api/v1/auth").Subrouter()

	api.HandleFunc("/login",
		h.login,
	).Methods(http.MethodPost)

	api.HandleFunc("/register",
		h.register,
	).Methods(http.MethodPost)
}
