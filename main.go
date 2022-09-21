package main

import (
	"log"
	"order/pkg/order/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	server := routes.Server{
		Router: r,
	}

	server.Register()

	log.Fatal(r.Run())
}
