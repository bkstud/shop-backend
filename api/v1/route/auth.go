package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func addAuthRoutes(rg *gin.RouterGroup) {
	loginGroup := rg.Group("/login")
	authGroup := rg.Group("/auth")

	loginGroup.GET("/google", func(c *gin.Context) {
		c.JSON(http.StatusOK, "login google")
	})

	loginGroup.GET("/github", func(c *gin.Context) {
		c.JSON(http.StatusOK, "login github")
	})

	authGroup.GET("/google", func(c *gin.Context) {
		c.JSON(http.StatusOK, "auth google")
	})

	authGroup.GET("/github", func(c *gin.Context) {
		c.JSON(http.StatusOK, "auth github")
	})
}
