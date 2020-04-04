// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/eddymoulton/stock-tracker/cmd/stocktracker/logger"
	"github.com/eddymoulton/stock-tracker/cmd/stocktracker/stocks"
	"github.com/eddymoulton/stock-tracker/cmd/stocktracker/transactions"
	"github.com/golobby/config"
	"github.com/jinzhu/gorm"
)

import (
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Injectors from wire.go:

func InitTransactionsAPI(db2 *gorm.DB, config2 *config.Config) transactions.API {
	loggerLogger := logger.ProvideLogger(config2)
	repository := transactions.ProvideTransactionsRepository(db2, loggerLogger)
	stocksRepository := stocks.ProvideStocksRepository(db2, loggerLogger)
	exchanges := stocks.ProvideExchanges(loggerLogger)
	service := stocks.ProvideStocksService(stocksRepository, exchanges, loggerLogger)
	transactionsService := transactions.ProvideTransactionsService(repository, service)
	api := transactions.ProvideTransactionsAPI(transactionsService)
	return api
}

func InitStocksAPI(db2 *gorm.DB, config2 *config.Config) stocks.API {
	loggerLogger := logger.ProvideLogger(config2)
	repository := stocks.ProvideStocksRepository(db2, loggerLogger)
	exchanges := stocks.ProvideExchanges(loggerLogger)
	service := stocks.ProvideStocksService(repository, exchanges, loggerLogger)
	api := stocks.ProvideStocksAPI(service)
	return api
}

func InitStocksService(db2 *gorm.DB, config2 *config.Config) stocks.Service {
	loggerLogger := logger.ProvideLogger(config2)
	repository := stocks.ProvideStocksRepository(db2, loggerLogger)
	exchanges := stocks.ProvideExchanges(loggerLogger)
	service := stocks.ProvideStocksService(repository, exchanges, loggerLogger)
	return service
}
