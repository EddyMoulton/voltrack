package transactions

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// API is a set of methods for managing transactions
type API struct {
	transactionsService *Service
}

// ProvideTransactionsAPI provides a new instance for wire
func ProvideTransactionsAPI(t *Service) *API {
	return &API{transactionsService: t}
}

// GetAll returns all transactions
func (a *API) GetAll(c *gin.Context) {
	transactions, err := a.transactionsService.GetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"transactions": transactions})
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
		a.transactionsService.AddBuyTransaction(data)
	} else if data.BuySell < 0 {
		a.transactionsService.AddSellTransaction(data)
	}

	c.JSON(http.StatusOK, gin.H{})
	return
}
