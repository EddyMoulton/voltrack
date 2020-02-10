package main

import (
	"log"

	"github.com/eddymoulton/stock-tracker/cmd/stocktracker/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// InitializeDatabase opens SQLite DB and migrates tables
func InitializeDatabase() *gorm.DB {
	db, err := gorm.Open("sqlite3", "stocktracker.db")

	if err != nil {
		log.Fatal(err)
		log.Fatal("Could not connect database")
	}

	db.AutoMigrate(&models.StockTransaction{}, &models.Transaction{})

	return db
}
