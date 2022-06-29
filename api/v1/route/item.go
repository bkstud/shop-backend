package route

import (
	"shop/api/v1/controller"
	"shop/api/v1/middleware"

	"github.com/gin-gonic/gin"
)

func addItemRoutes(rg *gin.RouterGroup) {
	itemGroup := rg.Group("/items")
	itemGroup.GET("/", controller.GetItems)

	itemGroup.Use(middleware.AuthorizeRequest())
	{
		itemGroup.POST("/", controller.CreateItem)

		itemGroup.PATCH("/:id", controller.EditItem)

		itemGroup.DELETE("/:id", controller.DeleteItem)
	}

}
