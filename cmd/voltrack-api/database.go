package main

import (
	"fmt"
	"log"
	"os"

	"github.com/eddymoulton/voltrack/pkg/reporting"
	"github.com/eddymoulton/voltrack/pkg/stocks"
	"github.com/eddymoulton/voltrack/pkg/transactions"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// InitializeDatabase opens SQLite DB and migrates tables
func InitializeDatabase() *gorm.DB {
	dbHost, exists := os.LookupEnv("DB_HOST")
	if !exists {
		dbHost = "localhost"
	}

	dbPort, exists := os.LookupEnv("DB_PORT")
	if !exists {
		dbPort = "5432"
	}

	dbUsername, exists := os.LookupEnv("DB_USER")
	if !exists {
		dbUsername = "postgres"
	}

	dbName, exists := os.LookupEnv("DB_NAME")
	if !exists {
		dbName = "postgres"
	}

	dbPassword, exists := os.LookupEnv("DB_PASSWORD")
	if !exists {
		dbPassword = "password"
	}

	devMode := ""
	_, exists = os.LookupEnv("ENV_DEVELOPMENT")
	if exists {
		devMode = "sslmode=disable"
	}

	db, err := gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s %s", dbHost, dbPort, dbUsername, dbName, dbPassword, devMode))

	if err != nil {
		log.Fatal(err)
		log.Fatal("Could not connect database")
	}
	1231
	db.AutoMigrate(&transactions.StockTransaction{},
		&transactions.Transaction{},
		&stocks.Stock{},
		&stocks.StockLog{},
		&reporting.OwnedStockLog{})

	return db
}
