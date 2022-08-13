package database

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"shop/api/v1/model"
	"strconv"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	Database *gorm.DB = nil
	DB_NAME           = "shop.db"
)

func init() {
	db, err := gorm.Open(sqlite.Open(DB_NAME), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&model.Item{})
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Transaction{})
	db.AutoMigrate(&model.Basket{})
	db.AutoMigrate(&model.BasketEntry{})
	db.AutoMigrate(&model.Feedback{})
	db.AutoMigrate(&model.Token{})

	Database = db

	if os.Getenv("VAR_CREATE_TEST_ITEMS") == "true" {
		createTestData()
	}
}

func createTestData() {
	for i := range [9]int{} {
		name := fmt.Sprintf("Product %c%d", 'A'+i, i)
		price, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", 100+rand.Float64()*150), 32)
		fprice := float32(price)
		item := model.Item{
			Name:        name,
			Description: "The cool description of '" + name + "'",
			Status:      "available",
			Price:       fprice,
		}
		if err := Database.Create(&item).Error; err != nil {
			log.Panic("Failed to create test item ", err)
		}
	}
}
