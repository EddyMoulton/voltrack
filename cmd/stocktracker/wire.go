//+build wireinject

package main

import (
	"github.com/eddymoulton/stock-tracker/cmd/stocktracker/logger"
	"github.com/eddymoulton/stock-tracker/cmd/stocktracker/stocks"
	"github.com/eddymoulton/stock-tracker/cmd/stocktracker/transactions"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
)

func InitTransactionAPI(db *gorm.DB, logger *logger.Logger) transactions.API {
	wire.Build(transactions.ProvideTransactionsRepository,
		transactions.ProvideTransactionsService,
		transactions.ProvideTransactionsAPI,
		stocks.ProvideStocksRepository,
		stocks.ProvideStocksService)

	return transactions.API{}
}
