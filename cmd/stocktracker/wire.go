//+build wireinject

package main

import (
	"github.com/eddymoulton/stock-tracker/cmd/stocktracker/transactions"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
)

func InitTransactionAPI(db *gorm.DB) transactions.TransactionAPI {
	wire.Build(transactions.ProvideTransactionRepository, transactions.ProvideTransactionService, transactions.ProvideTransactionAPI)

	return transactions.TransactionAPI{}
}
