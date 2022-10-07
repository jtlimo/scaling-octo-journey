package main

import (
	"log"
	"net/http"
	"order/pkg/order/routes"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	server := routes.Server{
		Router: r,
	}

	server.Register()

	log.Fatal(http.ListenAndServe(":3000", r))
}
