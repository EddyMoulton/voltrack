package stocks

import (
	"fmt"

	"github.com/eddymoulton/stock-tracker/cmd/stocktracker/logger"
	"github.com/jinzhu/gorm"
)

// Repository is a set of methods for handling stocks database access
type Repository struct {
	db     *gorm.DB
	logger *logger.Logger
}

// ProvideStocksRepository provides a new instance for wire
func ProvideStocksRepository(db *gorm.DB, logger *logger.Logger) *Repository {
	return &Repository{db, logger}
}

func (r *Repository) logDbAccess(message string) {
	r.logger.LogTrace("[DB]", message)
}

// Stock
func (r *Repository) getAll() ([]Stock, error) {

	allStocks := []Stock{}

	if err := r.db.Find(&allStocks).Error; err != nil {
		r.logger.LogWarning(err.Error())
		return allStocks, err
	}

	return allStocks, nil
}

func (r *Repository) find(code string) (Stock, error) {
	r.logDbAccess(fmt.Sprintf("Finding stock: %s", code))

	stock := Stock{}
	if err := r.db.Where(&Stock{Code: code}).Find(&stock).Error; err != nil {
		r.logger.LogWarning(err.Error())
		return stock, err
	}

	return stock, nil
}

func (r *Repository) add(stock Stock) (Stock, error) {
	r.logDbAccess("Adding stock")

	if stock.Code == "" {
		errorMessage := "Cannot add stock with empty code"
		r.logger.LogFatal(errorMessage)
		return Stock{}, fmt.Errorf(errorMessage)
	}

	if err := r.db.Create(&stock).Error; err != nil {
		r.logger.LogFatal(err.Error())
		return Stock{}, err
	}

	return stock, nil
}

// StockLog
func (r *Repository) addStockLogs(logs []StockLog) error {
	r.logDbAccess(fmt.Sprintf("Adding %v StockLogs", len(logs)))

	tx := r.db.Begin()

	for _, log := range logs {
		if err := tx.Create(&log).Error; err != nil {
			r.logger.LogFatal(err.Error())
			r.logDbAccess(fmt.Sprintf("Failed adding log for %s, rolling back", log.StockCode))
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}
