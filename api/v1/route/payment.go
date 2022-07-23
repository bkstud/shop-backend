package route

import (
	"shop/api/v1/controller"
	"shop/api/v1/middleware"

	"github.com/gin-gonic/gin"
)

func AddPaymentRoutes(rg *gin.RouterGroup) {
	paymentGroup := rg.Group("/payment")

	paymentGroup.Use(middleware.AuthorizeRequest())
	{
		// paymentGroup.GET("/", controller.CreateCheckoutSession)
		paymentGroup.POST("/create-checkout-session", controller.CreateCheckoutSession)
	}

}
