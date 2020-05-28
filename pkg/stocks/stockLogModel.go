package stocks

import (
	"time"
)

// StockLog is for storing the value of a Stock at a certain point in time
type StockLog struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Date      time.Time  `gorm:"primary_key;auto_increment:false"`
	StockCode string     `gorm:"primary_key;auto_increment:false"`
	Stock     *Stock     `gorm:"ForeignKey:StockCode;AssociationForeignKey:Code"`
	Value     int64      // Dollars x10^4
}

func stockLogFromDto(dto StockLogDto) StockLog {
	return StockLog{Date: dto.Date, StockCode: dto.StockCode, Value: dto.Value}
}
