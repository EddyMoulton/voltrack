package main

import (
	"net/http"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/golobby/config"
	"github.com/golobby/config/feeder"
	"github.com/jinzhu/gorm"
	"github.com/marcsantiago/gocron"
)

var db *gorm.DB

func main() {
	// Database
	db = InitializeDatabase()
	defer db.Close()

	// Configuration
	config, err := config.New(config.Options{
		Feeder: feeder.Map{
			"logLevel": "Trace",
		},
	})
	if err != nil {
		panic(err)
	}

	// Service initialisation
	stocksService := InitStocksService(db, config)
	transactionsAPI := InitTransactionsAPI(db, config)
	stocksAPI := InitStocksAPI(db, config)

	// Schedule
	gocron.Every(1).Minute().Do(stocksService.LogStocks)

	// HTTP
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

	go gocron.Start()
	router.Run(":3000")
}
