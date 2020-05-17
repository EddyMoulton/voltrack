package stocks

import (
	"time"
)

// StockLogDto is for transfering the value of a Stock at a certain point in time
type StockLogDto struct {
	Date      time.Time `json:"date" binding:"required"`      // Date of recording
	StockCode string    `json:"stockCode" binding:"required"` // Code of the stock
	Value     int64     `json:"value" binding:"required"`     // Dollars x10^4
}
