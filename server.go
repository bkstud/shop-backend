package main

import (
	"github.com/gin-gonic/gin"
	"shop/api/v1/route"
)

var (
	router = gin.Default()
)

func main() {
	v1 := router.Group("/api/v1")
	route.AddRoutes(v1)
	router.Run(":5000")
}
