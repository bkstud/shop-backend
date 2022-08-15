package model

import "gorm.io/gorm"

type Basket struct {
	gorm.Model
	UserEmail string
	Items     []BasketEntry `gorm:"foreignkey:BasketID"`
}

type BasketEntry struct {
	gorm.Model
	BasketID int
	ItemID   uint
	Item
}
