package transactions

import (
	"strconv"
	"time"

	"github.com/eddymoulton/stock-tracker/cmd/stocktracker/logger"
	"github.com/jinzhu/gorm"
)

// Repository is a set of methods for handling transaction database access
type Repository struct {
	db     *gorm.DB
	logger *logger.Logger
}

// ProvideTransactionsRepository provides a new instance for wire
func ProvideTransactionsRepository(db *gorm.DB, logger *logger.Logger) *Repository {
	return &Repository{db, logger}
}

func (r *Repository) logDbAccess(message ...string) {
	message = append([]string{"[DB]"}, message...)
	r.logger.LogTrace(message...)
}

func (r *Repository) getAll() ([]StockTransaction, error) {
	r.logDbAccess("Getting all transactions")

	allTransactions := []StockTransaction{}
	if err := r.db.Preload("BuyTransaction").Find(&allTransactions).Error; err != nil {
		r.logger.LogFatal(err.Error())
		return allTransactions, err
	}

	return allTransactions, nil
}

func (r *Repository) getOldestUnsoldStockTransactions(code string, limit int) ([]StockTransaction, error) {
	r.logDbAccess("Getting last", strconv.FormatInt(int64(limit), 10), "records for stock code", code)

	transactions := []StockTransaction{}
	if err := r.db.
		Preload("BuyTransaction").
		Preload("SellTransaction").
		Limit(limit).
		Where("sell_transaction_id = ?", "0").
		Where("stock_code = ?", code).
		Order("created_at asc").
		Find(&transactions).Error; err != nil {

		r.logger.LogFatal(err.Error())
		return transactions, err
	}

	return transactions, nil
}

func (r *Repository) addTransactions(transactions []StockTransaction) error {
	r.logDbAccess("Adding transactions")

	tx := r.db.Begin()

	for _, transaction := range transactions {
		if err := tx.Create(&transaction).Error; err != nil {
			r.logger.LogFatal(err.Error())
			r.logDbAccess("Failed adding transaction, rolling back")
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

func (r *Repository) updateTransactions(transactions []StockTransaction) error {
	r.logDbAccess("Adding transactions")

	tx := r.db.Begin()

	for _, transaction := range transactions {
		if err := tx.Save(&transaction).Error; err != nil {
			r.logger.LogFatal(err.Error())
			r.logDbAccess("Failed adding transaction, rolling back")
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

// GetStockTransactionsExistingBetween provides all stock transactions bought before the end date and sold after the start date
func (r *Repository) GetStockTransactionsExistingBetween(start, end time.Time) ([]StockTransaction, error) {
	r.logDbAccess("Getting stocks that existed between", start.Format("2006-01-02 15:04:05"), "and", end.Format("2006-01-02 15:04:05"))

	transactions := []StockTransaction{}
	if err := r.db.
		Preload("BuyTransaction").
		Preload("SellTransaction").
		Where("buy_transaction_id IN (?)", r.db.Table("transactions").Select("id").Where("date < ?", end.Format("2006-01-02 15:04:05")).QueryExpr()).
		Where("sell_transaction_id IN (?)", r.db.Table("transactions").Select("id").Where("date > ?", start.Format("2006-01-02 15:04:05")).QueryExpr()).
		Find(&transactions).Error; err != nil {

		r.logger.LogFatal(err.Error())
		return transactions, err
	}

	return transactions, nil
}
