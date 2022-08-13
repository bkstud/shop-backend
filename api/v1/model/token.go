package model

import "gorm.io/gorm"

type Token struct {
	gorm.Model
	Bearer    string `gorm:"uniqueIndex;not null;default:null"`
	UserEmail string `gorm:"uniqueIndex;not null;default:null"`
}
