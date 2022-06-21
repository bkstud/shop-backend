package main

import (
	"fmt"
	"shop/api/v1/route"
	"shop/config"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func main() {
	v1 := router.Group("/api/v1")
	route.AddRoutes(v1)

	router.Run(fmt.Sprintf(":%d", config.SERVER_PORT))
}
