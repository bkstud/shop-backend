package model

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	// will not allow name to be empty string
	Name        string  `gorm:"not null;type:varchar(100);default:null"`
	Status      string  `gorm:"not null;default:null"`
	Price       float32 `gorm:"not null;default:null"`
	Description string
}
