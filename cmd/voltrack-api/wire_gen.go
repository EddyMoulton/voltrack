// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/eddymoulton/voltrack/pkg/logger"
	"github.com/eddymoulton/voltrack/pkg/reporting"
	"github.com/eddymoulton/voltrack/pkg/stocks"
	"github.com/eddymoulton/voltrack/pkg/transactions"
	"github.com/golobby/config"
	"github.com/jinzhu/gorm"
)

import (
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Injectors from wire.go:

func InitTransactionsAPI(db2 *gorm.DB, config2 *config.Config) *transactions.API {
	loggerLogger := logger.ProvideLogger(config2)
	repository := transactions.ProvideTransactionsRepository(db2, loggerLogger)
	stocksRepository := stocks.ProvideStocksRepository(db2, loggerLogger)
	exchanges := stocks.ProvideExchanges(loggerLogger)
	service := stocks.ProvideStocksService(stocksRepository, exchanges, loggerLogger)
	transactionsService := transactions.ProvideTransactionsService(repository, service, loggerLogger)
	api := transactions.ProvideTransactionsAPI(transactionsService, loggerLogger)
	return api
}

func InitStocksAPI(db2 *gorm.DB, config2 *config.Config) *stocks.API {
	loggerLogger := logger.ProvideLogger(config2)
	repository := stocks.ProvideStocksRepository(db2, loggerLogger)
	exchanges := stocks.ProvideExchanges(loggerLogger)
	service := stocks.ProvideStocksService(repository, exchanges, loggerLogger)
	api := stocks.ProvideStocksAPI(service)
	return api
}

func InitStocksService(db2 *gorm.DB, config2 *config.Config) *stocks.Service {
	loggerLogger := logger.ProvideLogger(config2)
	repository := stocks.ProvideStocksRepository(db2, loggerLogger)
	exchanges := stocks.ProvideExchanges(loggerLogger)
	service := stocks.ProvideStocksService(repository, exchanges, loggerLogger)
	return service
}

func InitReportingAPI(db2 *gorm.DB, config2 *config.Config) *reporting.API {
	loggerLogger := logger.ProvideLogger(config2)
	repository := reporting.ProvideReportingRepository(db2, loggerLogger)
	transactionsRepository := transactions.ProvideTransactionsRepository(db2, loggerLogger)
	stocksRepository := stocks.ProvideStocksRepository(db2, loggerLogger)
	service := reporting.ProvideReportingService(repository, transactionsRepository, stocksRepository, loggerLogger)
	api := reporting.ProvideReportingAPI(service, loggerLogger)
	return api
}
