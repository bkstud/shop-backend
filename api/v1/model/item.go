package model

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	// will not allow name to be empty string
	Name        string `gorm:"not null;type:varchar(100);default:null"`
	Description string
	// TODO Find a way to limit possibilites to available, sold
	Status string
}
