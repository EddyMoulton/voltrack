package stocks

import "github.com/jinzhu/gorm"

type StockLog struct {
	gorm.Model
	StockCode string
	Stock     *Stock `gorm:"ForeignKey:StockCode;AssociationForeignKey:Code"`
	Value     int64  // Dollars x10^4
}
