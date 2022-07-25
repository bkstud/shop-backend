package controller

import (
	"net/http"
	"shop/api/v1/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

// GET /basket
// returns basket associated to user given by url params
func GetBasket(c *gin.Context) {
	email := c.DefaultQuery("email", "")
	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email not provided to query args."})
		return
	}
	var Basket model.Basket
	Db.Preload(clause.Associations).Find(&Basket, "WHERE UserEmail == ")
	c.JSON(http.StatusOK, Basket)
}

// Creates new basket for given user
func CreateBasket(c *gin.Context) {

}

// Updates user specific basket
func UpdateBasket(c *gin.Context) {

}

// Deletes user specific basket
func DeleteBasket(c *gin.Context) {

}
