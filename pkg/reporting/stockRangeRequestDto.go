package reporting

import "time"

// StockRangeRequestDto is used to provide a list of codes and date range to an API
type StockRangeRequestDto struct {
	StockCodes []string  `json:"stockCodes" binding:"required"`
	Start      time.Time `json:"start" binding:"required"`
	End        time.Time `json:"end" binding:"required"`
}
