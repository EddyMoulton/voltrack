package reporting

import (
	"time"

	"github.com/jinzhu/gorm"
)

// OwnedStockLog is a reporting model that captures historical data about a particular stock on a date
type OwnedStockLog struct {
	gorm.Model
	Date            time.Time
	StockCode       string
	Quantity        int64 // Dollars x10^4
	IndividualValue int64 // Dollars x10^4
	TotalValue      int64 // Dollars x10^4
}
