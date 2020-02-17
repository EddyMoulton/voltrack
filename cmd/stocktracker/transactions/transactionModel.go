package transactions

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Transaction struct {
	gorm.Model
	Date time.Time `json:"date" binding:"required"`
	Cost int64     `json:"cost" binding:"required"`
	Fee  int64     `json:"fee" binding:"required"`
}
