package main

import (
	"fmt"
	"net/http"

	"github.com/eddymoulton/stock-tracker/cmd/stocktracker/dtos"
	"github.com/eddymoulton/stock-tracker/cmd/stocktracker/models"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func main() {
	db = InitializeDatabase()

	router := gin.Default()

	router.Use(static.Serve("/", static.LocalFile("./views", true)))

	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}

	api.POST("/stocks/transaction", StockTransactionHandler)
	api.GET("/stocks/transaction", GetTransactionHandler)

	router.Run(":3000")
}

// StockTransactionHandler is an api endpoint handler for adding a new transaction of a particular stock
func StockTransactionHandler(c *gin.Context) {
	var json dtos.TransactionDto

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transactions := json.Map()

	if json.BuySell == 1 {
		// Buy, create new transactions
		transactionModels := make([]models.StockTransaction, json.Quantity)

		for index, element := range transactions {
			transactionModels[index] = models.StockTransaction{BuyTransaction: &element, StockCode: json.StockCode}
		}

		for _, model := range transactionModels {
			fmt.Printf("%+v\n", model)
			fmt.Printf("%+v\n", model.BuyTransaction)
			db.Create(&model)
		}
	} else if json.BuySell == -1 {
		// Sell, load existing transactions and add
		existingTransactions := make([]models.StockTransaction, 0)
		db.Where(&models.StockTransaction{StockCode: json.StockCode, SellTransaction: nil}).Order("CreatedAt").Limit(json.Quantity).Find(&existingTransactions)

		for index, element := range existingTransactions {
			element.SellTransaction = &transactions[index]
			db.Save(&element)
		}
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{
		"message":     fmt.Sprintf("Adding transaction of %d shares of %s, at a cost of $%0.2f each", json.Quantity, json.StockCode, float64(json.Cost)/10000),
		"tranactions": transactions,
	})
}

func GetTransactionHandler(c *gin.Context) {
	allTranactions := []models.StockTransaction{}
	db.Preload("BuyTransaction").Find(&allTranactions)

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{
		"tranactions": allTranactions,
	})
}
