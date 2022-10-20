package main

import (
	"log"
	"net/http"
	"order/pkg/order/application"
	"order/pkg/order/repository"
	"order/pkg/order/routes"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	repository := repository.New()
	application := application.New(repository)
	server := routes.Server{
		Router:      r,
		Application: application,
	}

	server.Register()

	log.Fatal(http.ListenAndServe(":3000", r))
}
