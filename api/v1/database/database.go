package database

import (
	"shop/api/v1/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Database *gorm.DB = nil

func init() {

	db, err := gorm.Open(sqlite.Open("shop.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&model.Item{})
	db.AutoMigrate(&model.Transaction{})

	Database = db
}
