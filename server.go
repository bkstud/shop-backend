package main

import (
	"shop/api/v1/route"

	"github.com/gin-gonic/gin"
)

var (
	routerHttps = gin.Default()
)

func main() {
	v1 := routerHttps.Group("/api/v1")
	route.AddRoutes(v1)
	routerHttps.RunTLS(":5000", "./cert/cert.pem", "./cert/key.pem")

}
