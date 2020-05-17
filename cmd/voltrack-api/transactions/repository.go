package transactions

import (
	"strconv"
	"time"

	"github.com/eddymoulton/stock-tracker/cmd/voltrack-api/logger"
	"github.com/jinzhu/gorm"
)

// Repository is a set of methods for handling transaction database access
type Repository struct {
	db  *gorm.DB
	log *logger.Logger
}

// ProvideTransactionsRepository provides a new instance for wire
func ProvideTransactionsRepository(db *gorm.DB, logger *logger.Logger) *Repository {
	return &Repository{db, logger}
}

func (r *Repository) getAll() ([]StockTransaction, error) {
	r.log.DbAccess("Getting all transactions")

	allTransactions := []StockTransaction{}
	if err := r.db.Preload("BuyTransaction").Find(&allTransactions).Error; err != nil {
		r.log.Error(err.Error())
		return allTransactions, err
	}

	return allTransactions, nil
}

func (r *Repository) getAllUnsoldStockTransactions() ([]StockTransaction, error) {
	r.log.DbAccess("Getting transactions for all unsold stocks")

	transactions := []StockTransaction{}
	if err := r.db.
		Preload("BuyTransaction").
		Preload("SellTransaction").
		Where("sell_transaction_id = ?", "0").
		Order("created_at asc").
		Find(&transactions).Error; err != nil {

		r.log.Error(err.Error())
		return transactions, err
	}

	r.log.Trace("Found", strconv.FormatInt(int64(len(transactions)), 10), "transactions")

	return transactions, nil
}

func (r *Repository) getOldestUnsoldStockTransactions(code string, limit int) ([]StockTransaction, error) {
	r.log.DbAccess("Getting last", strconv.FormatInt(int64(limit), 10), "records for stock code", code)

	transactions := []StockTransaction{}
	if err := r.db.
		Preload("BuyTransaction").
		Preload("SellTransaction").
		Limit(limit).
		Where("sell_transaction_id = ?", "0").
		Where("stock_code = ?", code).
		Order("created_at asc").
		Find(&transactions).Error; err != nil {

		r.log.Error(err.Error())
		return transactions, err
	}

	return transactions, nil
}

func (r *Repository) addTransactions(transactions []StockTransaction) error {
	r.log.DbAccess("Adding transactions")

	tx := r.db.Begin()

	for _, transaction := range transactions {
		if err := tx.Create(&transaction).Error; err != nil {
			r.log.Error(err.Error())
			r.log.DbAccess("Failed adding transaction, rolling back")
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

func (r *Repository) updateTransactions(transactions []StockTransaction) error {
	r.log.DbAccess("Adding transactions")

	tx := r.db.Begin()

	for _, transaction := range transactions {
		if err := tx.Save(&transaction).Error; err != nil {
			r.log.Error(err.Error())
			r.log.DbAccess("Failed adding transaction, rolling back")
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

// GetStockTransactionsExistingBetween provides all stock transactions bought before the end date and sold after the start date
func (r *Repository) GetStockTransactionsExistingBetween(start, end time.Time) ([]StockTransaction, error) {
	r.log.DbAccess("Getting stocks that existed between", start.Format("2006-01-02 15:04:05"), "and", end.Format("2006-01-02 15:04:05"))

	transactions := []StockTransaction{}
	if err := r.db.
		Preload("BuyTransaction").
		Preload("SellTransaction").
		Where("buy_transaction_id IN (?)", r.db.Table("transactions").Select("id").Where("date < ?", end.Format("2006-01-02 15:04:05")).QueryExpr()).
		Where("sell_transaction_id IN (?)", r.db.Table("transactions").Select("id").Where("date > ?", start.Format("2006-01-02 15:04:05")).QueryExpr()).
		Find(&transactions).Error; err != nil {

		r.log.Error(err.Error())
		return transactions, err
	}

	return transactions, nil
}
