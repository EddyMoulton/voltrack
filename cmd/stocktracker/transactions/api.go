package transactions

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type TransactionAPI struct {
	transactionService TransactionService
}

func ProvideTransactionAPI(t TransactionService) TransactionAPI {
	return TransactionAPI{transactionService: t}
}

func (api *TransactionAPI) GetAll(c *gin.Context) {
	transactions := api.transactionService.GetAll()

	c.JSON(http.StatusOK, gin.H{"transactions": transactions})
}

// StockTransactionHandler is an api endpoint handler for adding a new transaction of a particular stock
func AddHandler(c *gin.Context) {
	db, ok := c.MustGet("db").(gorm.DB)
	if !ok {
		// handle error here...
	}

	var json TransactionDto

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transactions := json.Map()

	if json.BuySell == 1 {
		// Buy, create new transactions
		transactionModels := make([]StockTransaction, json.Quantity)

		for index, element := range transactions {
			transactionModels[index] = StockTransaction{BuyTransaction: &element, StockCode: json.StockCode}
		}

		for _, model := range transactionModels {
			fmt.Printf("%+v\n", model)
			fmt.Printf("%+v\n", model.BuyTransaction)
			db.Create(&model)
		}
	} else if json.BuySell == -1 {
		// Sell, load existing transactions and add
		existingTransactions := make([]StockTransaction, 0)
		db.Where(&StockTransaction{StockCode: json.StockCode, SellTransaction: nil}).Order("CreatedAt").Limit(json.Quantity).Find(&existingTransactions)

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

func GetAllHandler(c *gin.Context) {
	db, ok := c.MustGet("db").(*gorm.DB)
	if !ok {
		fmt.Printf("Error in GetAllHandler %v\n", ok)
	}

	fmt.Printf("*** DB: %T: &i=%p\n", db, &db)

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{
		"tranactions": "", //allTranactions,
	})
}
