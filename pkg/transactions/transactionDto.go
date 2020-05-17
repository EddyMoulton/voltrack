package transactions

import (
	"time"
)

// TransactionDTO is used when adding a new purchase record
type TransactionDTO struct {
	BuySell   int64     `json:"buySell" binding:"required"`   // 1 for buy, -1 for sell
	Date      time.Time `json:"date" binding:"required"`      // When the transaction took place
	Cost      int64     `json:"cost" binding:"required"`      // Total cost of the transaction
	Fee       int64     `json:"fee" binding:"required"`       // Total fee paid
	StockCode string    `json:"stockCode" binding:"required"` // Short code of the stock
	Quantity  int       `json:"quantity" binding:"required"`  // Number of stocks in this transaction
}

// Map converts a single TransactionDto to many Transaction
func (t *TransactionDTO) Map() (transactions []Transaction) {
	transaction := Transaction{Date: t.Date, Cost: t.Cost / int64(t.Quantity), Fee: t.Fee / int64(t.Quantity)}

	transactions = make([]Transaction, t.Quantity)

	for i := 0; i < t.Quantity; i++ {
		transactions[i] = transaction
	}

	return transactions
}
