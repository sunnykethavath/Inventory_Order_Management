package database

import (
	"log"
	"os"

	"github.com/sunnykethavath/Inventory_Order_Management/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Declare the variable for the database
var DB *gorm.DB

// ConnectDB connect to db
func ConnectDB() {
	var err error

	dsn := "host=localhost user=postgres password=root dbname=Trademarkia port=5432 sslmode=disable"

	// Connect to the DB and initialize the DB variable
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect database! \n", err.Error())
		os.Exit(2)
	}
	DB.AutoMigrate(&model.User{}, &model.Product{}, &model.Order{})

	log.Println("Connection Opened to Database")
}
