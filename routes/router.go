package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"wagie.com/wageslavery/middlewares"
	foorouter "wagie.com/wageslavery/routes/foo_routes"
)

func InitRouter(r *mux.Router) {
	r.HandleFunc("/", indexHandler).Methods("GET")
	r.Use(middlewares.LoggingMiddleware)
	foorouter.CreateFooRouter(r)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "lets gooo!!!!")
}
