package controller

import (
	"errors"
	"fmt"
	"net/http"
	"shop/api/v1/model"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func findOrCreateBasketByEmail(c *gin.Context) *model.Basket {
	session := sessions.Default(c)
	email := session.Get("user-id")
	if email == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email not found is there a session?."})
		return nil
	}
	basket := new(model.Basket)
	result := Db.Model(&model.Basket{}).Preload("Items").First(&basket, "user_email = ?", email)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		basket.UserEmail = fmt.Sprintf("%v", email)
		basket.Items = []model.BasketEntry{}
		if err := Db.Create(&basket).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error()})
			return nil
		}

	}
	return basket
}

// GET /basket
// returns basket items associated to user logged in session
// If basket associated with user wasn't already created it will be
func GetBasketItems(c *gin.Context) {
	basket := findOrCreateBasketByEmail(c)
	if basket == nil {
		return
	}
	c.JSON(http.StatusOK, basket.Items)
}

// Creates new basket
func CreateBasket(c *gin.Context) {
	basket := new(model.Basket)
	if err := c.Bind(basket); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if err := Db.Create(basket).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, basket)
}

// Updates basket items
func UpdateBasketItems(c *gin.Context) {
	basket := findOrCreateBasketByEmail(c)
	if basket == nil {
		return
	}
	Db.Model(&basket).Association("Items")
	newItems := new([]model.Item)
	if err := c.BindJSON(newItems); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	newEntries := []model.BasketEntry{}

	for _, item := range *newItems {
		newEntries = append(newEntries, model.BasketEntry{BasketID: int(basket.ID), Item: item})
	}

	if err := Db.Model(&basket).Association("Items").Replace(&newEntries); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err})
		return
	}

	c.JSON(http.StatusOK, basket.Items)

}

// Deletes basket for specific user
func DeleteBasket(c *gin.Context) {
	basket := findOrCreateBasketByEmail(c)
	if err := Db.Model(&basket).Association("Items").Delete(&basket); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err})
		return
	}
	c.JSON(http.StatusOK, basket)
}
