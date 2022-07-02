package controller

import (
	"errors"
	"fmt"
	"net/http"
	"shop/api/v1/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func findItemById(c *gin.Context, id string) *model.Item {
	item := new(model.Item)
	result := Db.First(&item, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound,
			gin.H{"error": fmt.Sprintf("Item with id '%s' not found", id)})
		return nil

	}
	return item
}

// GET /items
func GetItems(c *gin.Context) {
	var items []model.Item
	Db.Find(&items)
	c.JSON(http.StatusOK, items)
}

// POST /items
func CreateItem(c *gin.Context) {
	item := new(model.Item)
	if err := c.Bind(item); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := Db.Create(item).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, item)
}

// PATCH /items/:id
func EditItem(c *gin.Context) {
	id := c.Param("id")
	item := findItemById(c, id)
	if item == nil {
		return
	}
	newItem := new(model.Item)
	if err := c.Bind(newItem); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := Db.Model(&item).Updates(newItem).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, item)
}

// DELETE /items/:id
func DeleteItem(c *gin.Context) {
	id := c.Param("id")
	item := findItemById(c, id)
	if item == nil {
		// The response is already handled by findItemById
		return
	}

	if err := Db.Delete(&item).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, item)
}
