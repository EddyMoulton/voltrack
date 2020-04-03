package main

import (
	"net/http"

	"github.com/eddymoulton/stock-tracker/cmd/stocktracker/logger"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func main() {
	db = InitializeDatabase()
	defer db.Close()

	logger := logger.Logger{}

	transactionsAPI := InitTransactionsAPI(db, &logger)
	stocksAPI := InitStocksAPI(db, &logger)

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

	api.GET("/stocks/transactions", transactionsAPI.GetAll)
	api.GET("/stocks", stocksAPI.GetAll)

	router.Run(":3000")
}
