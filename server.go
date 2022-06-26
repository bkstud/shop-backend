package main

import (
	"log"
	"shop/api/v1/auth"
	"shop/api/v1/route"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

var (
	routerHttps = gin.Default()
)

func main() {

	token, err := auth.RandToken(64)
	if err != nil {
		log.Fatal("unable to generate random token: ", err)
	}
	store := sessions.NewCookieStore([]byte(token))
	store.Options(sessions.Options{
		Path:   "/",
		MaxAge: 86400 * 7,
	})
	routerHttps.Use(gin.Logger())
	routerHttps.Use(gin.Recovery())
	routerHttps.Use(sessions.Sessions("store", store))

	v1 := routerHttps.Group("/api/v1")
	route.AddRoutes(v1)

	routerHttps.RunTLS(":5000", "./cert/cert.pem", "./cert/key.pem")

}
