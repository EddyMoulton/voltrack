package stocks

import (
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
	allStocks := []Stock{}
	if err := r.db.Find(&allStocks).Error; err != nil {
		r.logger.Log(err.Error())
		return allStocks, err
	}

	return allStocks, nil
}

func (r *Repository) find(code string) (Stock, error) {
	stock := Stock{}
	if err := r.db.Where(&Stock{Code: code}).Find(&stock).Error; err != nil {
		r.logger.Log(err.Error())
		return stock, err
	}

	return stock, nil
}

func (r *Repository) add(stock Stock) error {
	if err := r.db.Create(&stock).Error; err != nil {
		r.logger.Log(err.Error())
		return err
	}

	return nil
}
