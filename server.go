package main

import (
	"github.com/gin-gonic/gin"
	"shop/api/v1/route"
)

var (
	router = gin.Default()
)

// Run will start the server
func main() {
	v1 := router.Group("/api/v1")
	route.AddRoutes(v1)
	router.Run(":5000")
}
