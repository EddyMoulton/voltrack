package transactions

import (
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

func (r *Repository) logDbAccess(message string) {
	r.logger.LogTrace("[DB]", message)
}

func (r *Repository) getAll() []StockTransaction {
	r.logDbAccess("Getting all transactions")

	allTransactions := []StockTransaction{}
	r.db.Preload("BuyTransaction").Find(&allTransactions)
	return allTransactions
}

func (r *Repository) addTransactions(transactions []StockTransaction) {
	r.logDbAccess("Adding transactions")

	err := r.db.Transaction(func(tx *gorm.DB) error {
		for _, transaction := range transactions {
			if err := tx.Create(&transaction).Error; err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		r.logger.LogFatal(err.Error())
	}
}
