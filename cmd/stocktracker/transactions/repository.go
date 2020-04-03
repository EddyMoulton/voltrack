package transactions

import (
	"github.com/eddymoulton/stock-tracker/cmd/stocktracker/logger"
	"github.com/jinzhu/gorm"
)

// TransactionRepository is a set of methods for handling transaction database access
type TransactionRepository struct {
	db     *gorm.DB
	logger *logger.Logger
}

// ProvideTransactionRepository provides a new instance for wire
func ProvideTransactionRepository(db *gorm.DB, logger *logger.Logger) TransactionRepository {
	return TransactionRepository{db, logger}
}

func (r *TransactionRepository) getAll() []StockTransaction {
	allTransactions := []StockTransaction{}
	r.db.Preload("BuyTransaction").Find(&allTransactions)
	return allTransactions
}

func (r *TransactionRepository) addTransactions(transactions []StockTransaction) {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		for _, transaction := range transactions {
			if err := tx.Create(&transaction).Error; err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		r.logger.Log(err.Error())
	}
}
