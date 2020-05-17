package reporting

import "time"

// DateRangeDTO is used when a date range to an API
type DateRangeDTO struct {
	Start time.Time `json:"start" binding:"required"`
	End   time.Time `json:"end" binding:"required"`
}
