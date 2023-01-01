package foorouter

import (
	"github.com/gorilla/mux"
)

func CreateFooRouter(r *mux.Router) {
	apifoo := r.PathPrefix("/foo").Subrouter()
	SetAPIRoutes(apifoo)
}
