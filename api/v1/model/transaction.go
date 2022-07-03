package model

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	ItemID  uint
	Item    Item
	UserID  uint `gorm:"not null;type:varchar(100);default:null"`
	User    User
	Payment string // TODO: To be updated when stripe is added
	Type    string // purchase or return
	Status  string // pending or finished
}
