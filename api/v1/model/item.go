package model

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	Name        string
	Description string
	//TODO Find a way to limit possibilites to available, sold
	Status string
}
