package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"wagie.com/wageslavery/resources"
	router "wagie.com/wageslavery/routes"
)

func main() {
	resources.InitializeSQLite()
	defer resources.CloseSQLite()

	r := mux.NewRouter()
	router.InitRouter(r)

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
	}
	log.Printf("listening on %s", srv.Addr)
	log.Fatal(srv.ListenAndServe())

}
