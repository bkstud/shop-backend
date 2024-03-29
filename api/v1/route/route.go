package route

import (
	"github.com/gin-gonic/gin"
)

// Routes for /api/v1/ endpoint
func AddRoutes(rg *gin.RouterGroup) {
	addAuthRoutes(rg)
	addItemRoutes(rg)
	addTransactionRoutes(rg)
	addUserRoutes(rg)
	addPaymentRoutes(rg)
	addBasketRoutes(rg)
	addFeedbackRoutes(rg)
}
