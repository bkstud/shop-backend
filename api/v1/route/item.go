package route

import (
	"shop/api/v1/controller"

	"github.com/gin-gonic/gin"
)

func addItemRoutes(rg *gin.RouterGroup) {
	itemGroup := rg.Group("/items")

	itemGroup.GET("/", controller.GetItems)

	itemGroup.POST("/", controller.CreateItem)

	itemGroup.PATCH("/:id", controller.EditItem)

	itemGroup.DELETE("/:id", controller.DeleteItem)
}
