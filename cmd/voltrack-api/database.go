package main

import (
	"log"

	"github.com/eddymoulton/voltrack/pkg/reporting"
	"github.com/eddymoulton/voltrack/pkg/stocks"
	"github.com/eddymoulton/voltrack/pkg/transactions"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// InitializeDatabase opens SQLite DB and migrates tables
func InitializeDatabase() *gorm.DB {
	db, err := gorm.Open("mysql", "test_user:password@(10.1.1.11)/voltrack_dev?charset=utf8&parseTime=True&loc=Local")

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
