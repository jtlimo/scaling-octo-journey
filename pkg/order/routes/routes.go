package routes

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	Router *gin.Engine
}

func (s *Server) Register() {
	s.Router.POST("/order", createOrder)
}

func createOrder(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
