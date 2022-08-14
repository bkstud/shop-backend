package route

import (
	"shop/api/v1/controller"
	"shop/api/v1/middleware"

	"github.com/gin-gonic/gin"
)

func addPaymentRoutes(rg *gin.RouterGroup) {
	paymentGroup := rg.Group("/payment")

	paymentGroup.GET("/success", controller.HandleSuccess)
	paymentGroup.Use(middleware.AuthorizeRequest())
	{
		paymentGroup.POST("/create-checkout-session", controller.CreateCheckoutSession)
	}

}
