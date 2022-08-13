package middleware

import (
	"net/http"

	"shop/api/v1/controller"
	"shop/config"

	"github.com/gin-gonic/gin"
)

// AuthorizeRequest is used to authorize a request for a certain end-point group.
func AuthorizeRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		if config.ENV != "TEST" {

			bearer := c.GetHeader("Token")

			if bearer == "" {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Not logged in.",
					"message": "Please provide bearer token header."})
				c.Abort()
				return
			}

			token, exists := controller.GetToken(bearer)
			if !exists {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Not logged in.",
					"message": "Invalid token."})
				c.Abort()
				return

			}
			c.Set("user-email", token.UserEmail)
		}
		c.Next()
	}
}
