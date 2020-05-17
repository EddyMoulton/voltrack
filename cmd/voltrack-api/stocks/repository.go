package stocks

import (
	"fmt"
	"time"

	"github.com/eddymoulton/stock-tracker/cmd/voltrack-api/helpers"
	"github.com/eddymoulton/stock-tracker/cmd/voltrack-api/logger"
	"github.com/jinzhu/gorm"
)

// Repository is a set of methods for handling stocks database access
type Repository struct {
	db  *gorm.DB
	log *logger.Logger
}

// ProvideStocksRepository provides a new instance for wire
func ProvideStocksRepository(db *gorm.DB, logger *logger.Logger) *Repository {
	return &Repository{db, logger}
}

// Stock
func (r *Repository) getAll() ([]Stock, error) {
	allStocks := []Stock{}

	if err := r.db.Find(&allStocks).Error; err != nil {
		r.log.Warning(err.Error())
		return allStocks, err
	}

	return allStocks, nil
}

func (r *Repository) getAllCodes() ([]string, error) {
	allStocks := []Stock{}

	if err := r.db.Select([]string{"code"}).Find(&allStocks).Error; err != nil {
		r.log.Warning(err.Error())
		return []string{}, err
	}

	allStockCodes := make([]string, len(allStocks))
	for i := range allStocks {
		allStockCodes[i] = allStocks[i].Code
	}

	return allStockCodes, nil
}

func (r *Repository) find(code string) (Stock, error) {
	r.log.DbAccess(fmt.Sprintf("Finding stock: %s", code))

	stock := Stock{}
	if err := r.db.Where(&Stock{Code: code}).Find(&stock).Error; err != nil {
		r.log.Warning(err.Error())
		return stock, err
	}

	return stock, nil
}

func (r *Repository) add(stock Stock) (Stock, error) {
	r.log.DbAccess("Adding stock")

	if stock.Code == "" {
		errorMessage := "Cannot add stock with empty code"
		r.log.Error(errorMessage)
		return Stock{}, fmt.Errorf(errorMessage)
	}

	if err := r.db.Create(&stock).Error; err != nil {
		r.log.Error(err.Error())
		return Stock{}, err
	}

	return stock, nil
}

// StockLog
func (r *Repository) addStockLogs(logs []StockLog) error {
	r.log.DbAccess(fmt.Sprintf("Adding %v StockLogs", len(logs)))

	tx := r.db.Begin()

	for _, log := range logs {
		if err := tx.Create(&log).Error; err != nil {
			r.log.Error(err.Error())
			r.log.DbAccess(fmt.Sprintf("Failed adding log for %s, rolling back", log.StockCode))
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

func (r *Repository) GetStockLogs(stockCodes []string, start, end time.Time) ([]StockLog, error) {
	r.log.DbAccess("Getting StockLogs between",
		helpers.RemoveTime(start).Format("2006-01-02 15:04:05"),
		"and",
		helpers.RemoveTime(end).Add(24*time.Hour).Format("2006-01-02 15:04:05"))

	allStockLogs := []StockLog{}

	if err := r.db.
		Where("stock_code IN (?)", stockCodes).
		Where("date >= ?", helpers.RemoveTime(start)).
		Where("date <= ?", helpers.RemoveTime(end).Add(24*time.Hour)).
		Find(&allStockLogs).Error; err != nil {

		r.log.Warning(err.Error())
		return allStockLogs, err
	}

	return allStockLogs, nil
}

func (r *Repository) GetLatestStockLog(stockCode string) (StockLog, error) {
	r.log.DbAccess("Get last stock log for", stockCode)

	log := StockLog{}

	if err := r.db.
		Where("stock_code = ?", stockCode).
		Order("date desc").
		Last(&log).Error; err != nil {

		r.log.Warning(err.Error())
		log.StockCode = stockCode
		return log, nil // do not return error
	}

	return log, nil
}
