package route

import (
	"net/http"
	"shop/api/v1/middleware"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// GET api/v1/user
func addUserRoutes(rg *gin.RouterGroup) {
	userGroup := rg.Group("/user")
	userGroup.Use(middleware.AuthorizeRequest())
	{
		userGroup.GET("/", func(c *gin.Context) {
			session := sessions.Default(c)
			userEmail := session.Get("user-id")
			c.JSON(http.StatusOK, gin.H{
				"message": "User logged in",
				"user":    userEmail,
			})
		})

		userGroup.POST("/logout", func(c *gin.Context) {
			session := sessions.Default(c)
			session.Clear()
			session.Save()
		})
	}

}
