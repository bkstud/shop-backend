package controller

import (
	"fmt"
	"net/http"
	"shop/api/v1/model"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

// For handling GET /feedback
// Provides feedback that user with current session provided.
func GetFeedback(c *gin.Context) {
	var feedbacks []model.Feedback
	session := sessions.Default(c)
	email := session.Get("user-id")
	if email == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Email not found - anonymous feedback is reserved for admin user."})
		return
	}
	Db.Where("user_email = ?", email).Find(&feedbacks)
	c.JSON(http.StatusOK, feedbacks)
}

// For handling POST /feedback
// Creates new feedback object.
// If there is active session email field will be filled
// otherwise there is new anonymous feedback.
func CreateFeedback(c *gin.Context) {
	feedback := new(model.Feedback)
	if err := c.Bind(feedback); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	session := sessions.Default(c)
	email := session.Get("user-id")
	if email != nil {
		feedback.UserEmail = fmt.Sprintf("%v", email)
	}
	if err := Db.Create(feedback).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, feedback)
}
