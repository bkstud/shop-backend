package route

import (
	"shop/api/v1/controller"
	"shop/api/v1/middleware"

	"github.com/gin-gonic/gin"
)

func addTransactionRoutes(rg *gin.RouterGroup) {
	transactionGroup := rg.Group("/transactions")

	transactionGroup.Use(middleware.AuthorizeRequest())
	{
		transactionGroup.GET("/", controller.GetTransactions)

		transactionGroup.POST("/", controller.CreateTransaction)

		transactionGroup.PATCH("/:id", controller.EditTransaction)

		transactionGroup.DELETE("/:id", controller.DeleteTransaction)
	}
}
