package route

import (
	"shop/api/v1/controller"
	"shop/api/v1/middleware"

	"github.com/gin-gonic/gin"
)

func addBasketRoutes(rg *gin.RouterGroup) {
	basketGroup := rg.Group("/basket")
	basketGroup.Use(middleware.AuthorizeRequest())
	{
		basketGroup.GET("/", controller.GetBasket)

		basketGroup.POST("/", controller.CreateBasket)

		basketGroup.PATCH("/", controller.UpdateBasket)

		basketGroup.DELETE("/", controller.DeleteBasket)
	}

}
