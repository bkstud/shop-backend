package controller

import (
	"shop/api/v1/model"
	"shop/api/v1/utils"
)

// Returns token by bearer
func GetToken(bearer string) (*model.Token, bool) {
	token := model.Token{}
	if err := Db.First(&token, "bearer = ?", bearer).Error; err != nil {
		return nil, false
	}
	return &token, true
}

// Creates new token with given email
func CreateOrUpdateToken(email string) (*model.Token, error) {
	token := new(model.Token)
	Db.First(&token, "user_email = ?", email)
	bearer, _ := utils.RandString(32)
	token.Bearer = bearer
	token.UserEmail = email
	if err := Db.Save(&token).Error; err != nil {
		return nil, err
	}
	return token, nil
}
