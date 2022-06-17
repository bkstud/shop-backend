package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Nick  string
	Email string
	Token string
}
