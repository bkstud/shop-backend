package route

import (
	"shop/api/v1/controller"
	"shop/api/v1/middleware"

	"github.com/gin-gonic/gin"
)

func addFeedbackRoutes(rg *gin.RouterGroup) {
	feedbackGroup := rg.Group("/feedback")

	feedbackGroup.POST("/", controller.CreateFeedback)

	feedbackGroup.Use(middleware.AuthorizeRequest())
	{
		feedbackGroup.GET("/", controller.GetFeedback)
	}
}
