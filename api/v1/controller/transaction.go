package controller

import (
	"errors"
	"fmt"
	"net/http"
	"shop/api/v1/model"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func findTransactionById(c *gin.Context, id string) *model.Transaction {
	transaction := new(model.Transaction)
	result := Db.First(&transaction, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound,
			gin.H{"error": fmt.Sprintf("Transaction with id '%s' not found", id)})
		return nil

	}
	return transaction
}

// GET /transactions
// returns all transactions for user session
func GetTransactions(c *gin.Context) {
	session := sessions.Default(c)
	email := session.Get("user-id")
	var user model.User
	Db.Where("email = ?", email).First(&user)
	var transactions []model.Transaction
	Db.Preload("User").Preload("Item").Where("user_id = ?", user.ID).Find(&transactions)
	c.JSON(http.StatusOK, transactions)
}

// POST /transactions
func CreateTransaction(c *gin.Context) {
	transaction := new(model.Transaction)
	if err := c.Bind(transaction); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if err := Db.Create(transaction).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, transaction)
}

// PATCH /transactions/:id
// edits transaction given by id
func EditTransaction(c *gin.Context) {
	id := c.Param("id")
	transaction := findTransactionById(c, id)
	if transaction == nil {
		return
	}
	newTransaction := new(model.Transaction)
	if err := c.Bind(newTransaction); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := Db.Model(&transaction).Updates(newTransaction).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, transaction)
}

// DELETE /transactions/:id
// delete trasaction with given id from db
func DeleteTransaction(c *gin.Context) {
	id := c.Param("id")
	transaction := findTransactionById(c, id)
	if transaction == nil {
		// The response is already handled by findTransactionById
		return
	}

	if err := Db.Delete(&transaction).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, transaction)
}
