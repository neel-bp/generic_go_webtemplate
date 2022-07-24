package foorouter

import (
	"github.com/gorilla/mux"
	"wagie.com/wageslavery/controllers/foo_controllers"
)

var controller = foo_controllers.New()

func SetAPIRoutes(r *mux.Router) {
	r.HandleFunc("", controller.Get).Methods("GET")
	r.HandleFunc("/", controller.Get).Methods("GET")
	r.HandleFunc("", controller.Create).Methods("POST")
	r.HandleFunc("/", controller.Create).Methods("POST")
}
