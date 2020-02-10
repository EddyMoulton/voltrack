package dtos

import (
	"time"

	"github.com/eddymoulton/stock-tracker/cmd/stocktracker/models"
)

// StockTransaction is used when adding a new purchase record
type TransactionDto struct {
	BuySell   int64     `json:"buySell" binding:"required"`
	Date      time.Time `json:"date" binding:"required"`
	Cost      int64     `json:"cost" binding:"required"`
	Fee       int64     `json:"fee" binding:"required"`
	StockCode string    `json:"stockCode" binding:"required"`
	Quantity  int       `json:"quantity" binding:"required"`
}

func (t *TransactionDto) Map() (transactions []models.Transaction) {
	transaction := models.Transaction{Date: t.Date, Cost: t.Cost, Fee: t.Fee / int64(t.Quantity)}

	transactions = make([]models.Transaction, t.Quantity)

	for i := 0; i < t.Quantity; i++ {
		transactions[i] = transaction
	}

	return
}
