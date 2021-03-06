package transactions

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Transaction is used to save data for a single buy/sell event (database will collate common data though)
type Transaction struct {
	gorm.Model
	Date time.Time `json:"date" binding:"required"`
	Cost int64     `json:"cost" binding:"required"`
	Fee  int64     `json:"fee" binding:"required"`
}
