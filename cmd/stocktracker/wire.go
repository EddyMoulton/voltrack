//+build wireinject

package main

import (
	"github.com/eddymoulton/stock-tracker/cmd/stocktracker/logger"
	"github.com/eddymoulton/stock-tracker/cmd/stocktracker/reporting"
	"github.com/eddymoulton/stock-tracker/cmd/stocktracker/stocks"
	"github.com/eddymoulton/stock-tracker/cmd/stocktracker/transactions"
	"github.com/golobby/config"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
)

func InitTransactionsAPI(db *gorm.DB, config *config.Config) *transactions.API {
	wire.Build(logger.ProvideLogger,
		transactions.ProvideTransactionsRepository,
		transactions.ProvideTransactionsService,
		transactions.ProvideTransactionsAPI,
		stocks.ProvideStocksRepository,
		stocks.ProvideStocksService,
		stocks.ProvideExchanges)

	return &transactions.API{}
}

func InitStocksAPI(db *gorm.DB, config *config.Config) *stocks.API {
	wire.Build(logger.ProvideLogger,
		stocks.ProvideStocksRepository,
		stocks.ProvideStocksService,
		stocks.ProvideExchanges,
		stocks.ProvideStocksAPI)

	return &stocks.API{}
}

func InitStocksService(db *gorm.DB, config *config.Config) *stocks.Service {
	wire.Build(logger.ProvideLogger,
		stocks.ProvideStocksRepository,
		stocks.ProvideStocksService,
		stocks.ProvideExchanges)

	return &stocks.Service{}
}

func InitReportingAPI(db *gorm.DB, config *config.Config) *reporting.API {
	wire.Build(logger.ProvideLogger,
		reporting.ProvideReportingRepository,
		reporting.ProvideReportingService,
		reporting.ProvideReportingAPI)

	return &reporting.API{}
}
