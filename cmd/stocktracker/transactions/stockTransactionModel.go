package transactions

import (
	"github.com/jinzhu/gorm"
)

// StockTransaction is used when adding a new purchase record
type StockTransaction struct {
	gorm.Model
	BuyTransaction  *Transaction `json:"buyTransaction" binding:"required"`
	SellTransaction *Transaction `json:"sellTransaction" binding:"required"`
	StockCode       string       `json:"stockCode" binding:"required"`
}

func (s *StockTransaction) BuyNew(code string, t *Transaction) {
	s.StockCode = code
	s.BuyTransaction = t
}

func (s *StockTransaction) Sell(t *Transaction) {
	s.SellTransaction = t
}
