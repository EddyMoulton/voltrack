package transactions

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// API is a set of methods for managing transactions
type API struct {
	service *Service
}

// ProvideTransactionsAPI provides a new instance for wire
func ProvideTransactionsAPI(t *Service) *API {
	return &API{service: t}
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

// AddTransaction creates a new set of transactions
func (a *API) AddTransaction(c *gin.Context) {
	var data TransactionDTO

	if err := c.ShouldBindJSON(&data); err != nil {
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
