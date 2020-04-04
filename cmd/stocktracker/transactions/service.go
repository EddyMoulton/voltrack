package transactions

import (
	"github.com/eddymoulton/stock-tracker/cmd/stocktracker/logger"
	"github.com/eddymoulton/stock-tracker/cmd/stocktracker/stocks"
)

// Service is an object that provides methods for altering or manipulating stock transactions
type Service struct {
	repository    *Repository
	stocksService *stocks.Service
	logger        *logger.Logger
}

// ProvideTransactionsService is a method to handle DI
func ProvideTransactionsService(r *Repository, s *stocks.Service, logger *logger.Logger) *Service {
	return &Service{r, s, logger}
}

// GetAll returns all the transactions in the database
func (s *Service) GetAll() ([]StockTransaction, error) {
	return s.repository.getAll()
}

// AddBuyTransaction adds a new set of transactions to the repositoryervice
func (s *Service) AddBuyTransaction(transactionDTO TransactionDTO) {
	stock, err := s.stocksService.Find(transactionDTO.StockCode)

	if err != nil {
		stock, err = s.stocksService.AddStock(transactionDTO.StockCode)
	}

	if err == nil {
		transactions := transactionDTO.Map()

		stockTransactions := make([]StockTransaction, len(transactions))
		for i, transaction := range transactions {
			stockTransaction := StockTransaction{}

			stockTransaction.BuyNew(&stock, &transaction)

			stockTransactions[i] = stockTransaction
		}

		s.repository.addTransactions(stockTransactions)
	}
}

// AddSellTransaction adds a new set of transactions to the repositoryervice
func (s *Service) AddSellTransaction(transactionDTO TransactionDTO) {
	transactions := transactionDTO.Map()

	stockTransactions, err := s.repository.getOldestUnsoldStockTransactions(transactionDTO.StockCode, len(transactions))

	if err != nil {
		s.logger.LogFatal("Cannot find stock to add sale transaction for")
	}

	for i, transaction := range transactions {
		stockTransactions[i].Sell(&transaction)
	}

	s.repository.updateTransactions(stockTransactions)
}
