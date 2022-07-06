package route

import (
	"github.com/gin-gonic/gin"
)

func AddRoutes(rg *gin.RouterGroup) {
	addAuthRoutes(rg)
	addItemRoutes(rg)
	addTransactionRoutes(rg)
	addUserRoutes(rg)
}
