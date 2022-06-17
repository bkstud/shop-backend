package controller

import (
	"errors"
	"fmt"
	"net/http"
	"shop/api/v1/database"
	"shop/api/v1/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db = database.Database

func findItemById(c *gin.Context, id string) *model.Item {
	item := new(model.Item)
	result := db.First(&item, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound,
			fmt.Sprintf("Item with id %s not found", id))
		return nil

	}
	return item
}

// GET /items
func GetItems(c *gin.Context) {
	var items []model.Item
	db.Find(&items)
	c.JSON(http.StatusOK, items)
}

// POST /items
func CreateItem(c *gin.Context) {
	item := new(model.Item)
	if err := c.Bind(item); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
	db.Create(item)
	c.JSON(http.StatusOK, item)
}

// PATCH /items/:id
func EditItem(c *gin.Context) {
	id := c.Param("id")
	item := findItemById(c, id)
	if item == nil {
		return
	}

	c.JSON(http.StatusOK, item)
}

// DELETE /items/:id
func DeleteItem(c *gin.Context) {
	id := c.Param("id")
	item := findItemById(c, id)
	if item == nil {
		return
	}
	db.Delete(&item)
	c.JSON(http.StatusOK, item)
}
