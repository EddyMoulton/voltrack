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
func (api *API) GetAll(c *gin.Context) {
	transactions := api.transactionsService.GetAll()

	c.JSON(http.StatusOK, gin.H{"transactions": transactions})
	return
}

// AddTransaction creates a new set of transactions
func (api *API) AddTransaction(c *gin.Context) {
	var data TransactionDTO

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	api.transactionsService.AddTransaction(data)

	c.JSON(http.StatusOK, gin.H{})
	return
}
