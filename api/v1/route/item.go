package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func addItemRoutes(rg *gin.RouterGroup) {
	itemGroup := rg.Group("/items")
	itemGroup.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "returns list of items")
	})

	itemGroup.POST("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "adds new item to db")
	})

	itemGroup.PATCH("/:id", func(c *gin.Context) {
		c.JSON(http.StatusOK, "edits item with id")
	})

	itemGroup.DELETE("/:id", func(c *gin.Context) {
		c.JSON(http.StatusOK, "deletes item with id")
	})

}
