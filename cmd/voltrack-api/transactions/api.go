package transactions

import (
	"net/http"

	"github.com/eddymoulton/stock-tracker/cmd/voltrack-api/logger"
	"github.com/gin-gonic/gin"
)

// API is a set of methods for managing transactions
type API struct {
	service *Service
	log     *logger.Logger
}

// ProvideTransactionsAPI provides a new instance for wire
func ProvideTransactionsAPI(t *Service, log *logger.Logger) *API {
	return &API{service: t, log: log}
}

// GetAll returns all transactions
func (a *API) GetAll(c *gin.Context) {
	transactions, err := a.service.GetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"transactions": transactions})
	return
}

// GetCurrentStocks returns all the stock objects in the database that are currently owned along with basic stas
func (a *API) GetCurrentStocks(c *gin.Context) {
	stocks, err := a.service.GetCurrentStocks()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"currentStocks": stocks})
	return
}

// GetTransactionSummaries returns all the transactions in the database that are currently owned along with their recent values
func (a *API) GetTransactionSummaries(c *gin.Context) {
	stocks, err := a.service.GetTransactionSummaries()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"transactions": stocks})
	return
}

// AddTransaction creates a new set of transactions
func (a *API) AddTransaction(c *gin.Context) {
	var data TransactionDTO

	if err := c.ShouldBindJSON(&data); err != nil {
		a.log.Warning(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if data.BuySell > 0 {
		a.service.AddBuyTransaction(data)
	} else if data.BuySell < 0 {
		a.service.AddSellTransaction(data)
	}

	c.JSON(http.StatusOK, gin.H{})
	return
}
