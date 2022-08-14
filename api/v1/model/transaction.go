package model

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	ItemID  uint
	Item    Item
	UserID  uint `gorm:"not null;type:varchar(100);default:null"`
	User    User
	Payment string
	// Type of transaction cant be: purchase or return
	Type string
	// The realization status - providing product to customer
	// Either pending or finished
	Status string
	// id of stripe session
	SessionID string `gorm:"not null;default:null"`
}
