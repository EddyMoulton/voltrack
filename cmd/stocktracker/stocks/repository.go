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
func ProvideStocksRepository(db *gorm.DB, logger *logger.Logger) Repository {
	return Repository{db, logger}
}

func (r *Repository) getAll() ([]Stock, error) {
	r.logger.LogTrace("[DB] Getting all stocks")

	allStocks := []Stock{}

	if err := r.db.Find(&allStocks).Error; err != nil {
		r.logger.LogError(err.Error())
		return allStocks, err
	}

	return allStocks, nil
}

func (r *Repository) find(code string) (Stock, error) {
	r.logger.LogTrace(fmt.Sprintf("[DB] Finding stock: %s", code))

	stock := Stock{}
	if err := r.db.Where(&Stock{Code: code}).Find(&stock).Error; err != nil {
		r.logger.LogError(err.Error())
		return stock, err
	}

	return stock, nil
}

func (r *Repository) add(stock Stock) error {
	r.logger.LogTrace("[DB] Adding stock")

	if err := r.db.Create(&stock).Error; err != nil {
		r.logger.LogError(err.Error())
		return err
	}

	return nil
}
