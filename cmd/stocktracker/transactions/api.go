package transactions

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// TransactionAPI is a set of methods for managing transactions
type TransactionAPI struct {
	transactionService TransactionService
}

// ProvideTransactionAPI provides a new instance for wire
func ProvideTransactionAPI(t TransactionService) TransactionAPI {
	return TransactionAPI{transactionService: t}
}

// GetAll returns all transactions
func (api *TransactionAPI) GetAll(c *gin.Context) {
	transactions := api.transactionService.GetAll()

	c.JSON(http.StatusOK, gin.H{"transactions": transactions})
	return
}

// AddTransaction creates a new set of transactions
func (api *TransactionAPI) AddTransaction(c *gin.Context) {
	var data TransactionDto

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	api.transactionService.AddTransaction(data)

	c.JSON(http.StatusOK, gin.H{})
	return
}
