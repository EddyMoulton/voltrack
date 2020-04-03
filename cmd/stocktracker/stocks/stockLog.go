package stocks

import "github.com/jinzhu/gorm"

// StockLog is for storing the value of a Stock at a certain point in time
type StockLog struct {
	gorm.Model
	StockCode string
	Stock     *Stock `gorm:"ForeignKey:StockCode;AssociationForeignKey:Code"`
	Value     int64  // Dollars x10^4
}
