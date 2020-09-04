package main

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
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
	reportingAPI := InitReportingAPI(db, config)

	// Schedule
	gocron.ChangeLoc(time.Now().UTC().Location())
	gocron.Every(1).Day().At("9:00").Do(stocksService.LogStocks)

	// HTTP
	router := gin.Default()
	router.OPTIONS("/*any", preflight)

	router.Use(static.Serve("/", static.LocalFile("./views", true)))
	router.Use(cors.Default())

	api := router.Group("/")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}

	api.GET("/transactions/summaries", transactionsAPI.GetTransactionSummaries)
	api.GET("/stocks/history", reportingAPI.GetOwnedStockLogs)
	api.PUT("/reporting/generate", reportingAPI.GenerateSummaryLogs)
	api.GET("/stocks/transactions", transactionsAPI.GetAll)
	api.POST("/stocks/transactions", transactionsAPI.AddTransaction)
	api.POST("/stocks/transactions/bulk", transactionsAPI.UploadTransactionHistory)
	api.GET("/stocks", stocksAPI.GetAll)
	api.GET("/stocks/current", transactionsAPI.GetCurrentStocks)
	api.GET("/stocks/logs", stocksAPI.GetStockHistory)
	api.POST("/stocks/logs", stocksAPI.UploadStockHistory)

	go gocron.Start()
	router.Run(":3000")
}

func preflight(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
	c.JSON(http.StatusOK, struct{}{})
}
