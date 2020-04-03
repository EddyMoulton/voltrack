package transactions

import (
	"github.com/eddymoulton/stock-tracker/cmd/stocktracker/stocks"
	"github.com/jinzhu/gorm"
)

// StockTransaction is used when adding a new purchase record
type StockTransaction struct {
	gorm.Model
	BuyTransactionID  int64
	BuyTransaction    *Transaction `gorm:"ForeignKey:BuyTransactionID;AssociationForeignKey:ID"`
	SellTransactionID int64
	SellTransaction   *Transaction `gorm:"ForeignKey:SellTransactionID;AssociationForeignKey:ID"`
	StockCode         string
	Stock             *stocks.Stock `gorm:"ForeignKey:StockCode;AssociationForeignKey:Code"`
}

// BuyNew creates a StockTranscation from a stock and transaction
func (s *StockTransaction) BuyNew(stock *stocks.Stock, t *Transaction) {
	s.Stock = stock
	s.BuyTransaction = t
}

// Sell modifies an existing StockTransaction with sale data
func (s *StockTransaction) Sell(t *Transaction) {
	s.SellTransaction = t
}
