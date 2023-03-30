package main

import (
	"log"
	"net/http"

	"github.com/dev-hack95/bookstore_api_backend/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStore(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("0.0.0.0:9010", r))
}
