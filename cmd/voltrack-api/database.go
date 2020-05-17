package main

import (
	"log"

	"github.com/eddymoulton/voltrack/pkg/reporting"
	"github.com/eddymoulton/voltrack/pkg/stocks"
	"github.com/eddymoulton/voltrack/pkg/transactions"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// InitializeDatabase opens SQLite DB and migrates tables
func InitializeDatabase() *gorm.DB {
	db, err := gorm.Open("sqlite3", "voltrack-api.db")

	if err != nil {
		log.Fatal(err)
		log.Fatal("Could not connect database")
	}

	db.AutoMigrate(&transactions.StockTransaction{},
		&transactions.Transaction{},
		&stocks.Stock{},
		&stocks.StockLog{},
		&reporting.OwnedStockLog{})

	return db
}
