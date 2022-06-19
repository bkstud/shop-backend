package database

import (
	"os"
	"shop/api/v1/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	Database *gorm.DB = nil
	DB_NAME           = "shop.db"
)

func init() {
	// Recreate every time as this will also happen in the container.
	// Do not use on production.
	if _, err := os.Stat(DB_NAME); err == nil {
		os.Remove(DB_NAME)
	}

	db, err := gorm.Open(sqlite.Open(DB_NAME), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&model.Item{})
	db.AutoMigrate(&model.Transaction{})

	Database = db
}
