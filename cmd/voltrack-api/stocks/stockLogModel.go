package stocks

import (
	"time"

	"github.com/jinzhu/gorm"
)

// StockLog is for storing the value of a Stock at a certain point in time
type StockLog struct {
	gorm.Model
	Date      time.Time
	StockCode string
	Stock     *Stock `gorm:"ForeignKey:StockCode;AssociationForeignKey:Code"`
	Value     int64  // Dollars x10^4
}

func stockLogFromDto(dto StockLogDto) StockLog {
	return StockLog{Date: dto.Date, StockCode: dto.StockCode, Value: dto.Value}
}
