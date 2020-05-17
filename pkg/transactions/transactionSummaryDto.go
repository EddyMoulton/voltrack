package transactions

import (
	"time"
)

// TransactionSummaryDTO is used when adding a new purchase record
type TransactionSummaryDTO struct {
	Code          string    `json:"code" binding:"required"`          // Short code of the stock
	Date          time.Time `json:"date" binding:"required"`          // When the transaction took place
	Cost          int64     `json:"cost" binding:"required"`          // Cost per share
	Value         int64     `json:"value" binding:"required"`         // Current value per share
	DividendValue int64     `json:"dividendValue" binding:"required"` // Total dividends earned per share
	Quantity      int       `json:"quantity" binding:"required"`      // Number of stocks in this transaction
}
