package transactions

import "github.com/jinzhu/gorm"

type TransactionRepository struct {
	db *gorm.DB
}

func ProvideTransactionRepository(db *gorm.DB) TransactionRepository {
	return TransactionRepository{db}
}

func (r *TransactionRepository) GetAll() []StockTransaction {
	allTranactions := []StockTransaction{}
	r.db.Preload("BuyTransaction").Find(&allTranactions)
	return allTranactions
}
