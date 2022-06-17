package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func addTransactionRoutes(rg *gin.RouterGroup) {
	transactionGroup := rg.Group("/transactions")

	transactionGroup.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "returns list of transactions")
	})

	transactionGroup.POST("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "adds new transaction to db")
	})

	transactionGroup.PATCH("/:id", func(c *gin.Context) {
		c.JSON(http.StatusOK, "edits transaction with id")
	})

	transactionGroup.DELETE("/:id", func(c *gin.Context) {
		c.JSON(http.StatusOK, "deletes transaction with id")
	})

}
