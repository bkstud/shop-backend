package main

import (
	"shop/api/v1/route"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func main() {
	v1 := router.Group("/api/v1")
	route.AddRoutes(v1)
	router.Run(":5000")
}
