package route

import (
	"shop/api/v1/auth/github"
	"shop/api/v1/auth/google"

	"github.com/gin-gonic/gin"
)

func addAuthRoutes(rg *gin.RouterGroup) {
	loginGroup := rg.Group("/login")
	authGroup := rg.Group("/auth")

	loginGroup.GET("/google", google.LoginHandler)
	loginGroup.GET("/github", github.LoginHandler)

	authGroup.GET("/google", google.AuthHandler)
	authGroup.GET("/github", github.AuthHandler)
}
