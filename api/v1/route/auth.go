package route

import (
	"shop/api/v1/auth/facebook"
	"shop/api/v1/auth/github"
	"shop/api/v1/auth/google"

	"github.com/gin-gonic/gin"
)

func addAuthRoutes(rg *gin.RouterGroup) {
	loginGroup := rg.Group("/login")
	authGroup := rg.Group("/auth")

	loginGroup.GET("/google", google.LoginHandler)
	loginGroup.GET("/github", github.LoginHandler)
	loginGroup.GET("/facebook", facebook.LoginHandler)

	authGroup.GET("/google", google.AuthHandler)
	authGroup.GET("/github", github.AuthHandler)
	authGroup.GET("/facebook", facebook.AuthHandler)
}
