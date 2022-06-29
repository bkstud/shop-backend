package middleware

import (
	"net/http"

	"shop/config"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

// AuthorizeRequest is used to authorize a request for a certain end-point group.
func AuthorizeRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		if config.ENV != "TEST" {
			session := sessions.Default(c)
			v := session.Get("user-id")
			if v == nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Not logged in.",
					"message": "Please log in."})
				c.Abort()
			}
		}
		c.Next()
	}
}
