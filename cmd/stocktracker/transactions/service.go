package transactions

import "github.com/eddymoulton/stock-tracker/cmd/stocktracker/stocks"

// Service is an object that provides methods for altering or manipulating stock transactions
type Service struct {
	repository    Repository
	stocksService stocks.Service
}

// ProvideTransactionsService is a method to handle DI
func ProvideTransactionsService(r Repository, s stocks.Service) Service {
	return Service{r, s}
}

// GetAll returns all the transactions in the database
func (service *Service) GetAll() []StockTransaction {
	return service.repository.getAll()
}

// AddTransaction adds a new set of transactions to the repositoryervice
func (service *Service) AddTransaction(transaction TransactionDTO) {
	stock, err := service.stocksService.Find(transaction.StockCode)

	if err != nil {
		transactions := transaction.Map()

		stockTransactions := make([]StockTransaction, len(transactions))
		for i, transaction := range transactions {
			stockTransaction := StockTransaction{}

			stockTransaction.BuyNew(&stock, &transaction)

			stockTransactions[i] = stockTransaction
		}

		service.repository.addTransactions(stockTransactions)
	}
}
