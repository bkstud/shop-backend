package model

import "gorm.io/gorm"

type Basket struct {
	gorm.Model
	UserEmail string `gorm:"not null;type:varchar(100);default:null"`
	User      User
	Items     []Item
}
