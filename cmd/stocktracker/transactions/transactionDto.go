package transactions

import (
	"time"
)

// TransactionDTO is used when adding a new purchase record
type TransactionDTO struct {
	BuySell   int64     `json:"buySell" binding:"required"`
	Date      time.Time `json:"date" binding:"required"`
	Cost      int64     `json:"cost" binding:"required"`
	Fee       int64     `json:"fee" binding:"required"`
	StockCode string    `json:"stockCode" binding:"required"`
	Quantity  int       `json:"quantity" binding:"required"`
}

// Map converts a single TransactionDto to many Transaction
func (t *TransactionDTO) Map() (transactions []Transaction) {
	transaction := Transaction{Date: t.Date, Cost: t.Cost, Fee: t.Fee / int64(t.Quantity)}

	transactions = make([]Transaction, t.Quantity)

	for i := 0; i < t.Quantity; i++ {
		transactions[i] = transaction
	}

	return transactions
}
