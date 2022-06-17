package model

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	ItemID  Item
	Item    Item
	BuyerID uint
	Buyer   User
	Payment string
}
