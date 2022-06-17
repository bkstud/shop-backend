package model

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	ItemID  uint
	Item    Item
	BuyerID uint
	Buyer   User
	Payment string
}
