package stocks

import "time"

// GetStockLogsDto contains data to control what stock logs are returned
type GetStockLogsDto struct {
	StockCodes []string  `json:"stockCodes" binding:"required"` // List of stock codes to get logs for
	Start      time.Time `json:"start"`
	End        time.Time `json:"end"`
}
