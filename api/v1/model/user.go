package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Identity string
	Token    string
}
