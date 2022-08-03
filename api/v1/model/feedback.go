package model

import "gorm.io/gorm"

type Feedback struct {
	gorm.Model
	UserEmail string
	Contents  string
	Response  string
}
