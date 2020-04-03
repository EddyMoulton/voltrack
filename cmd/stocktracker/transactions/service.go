package transactions

import "github.com/eddymoulton/stock-tracker/cmd/stocktracker/stocks"

// TransactionService is an object that provides methods for altering or manipulating stock transactions
type TransactionService struct {
	transactionRepository TransactionRepository
	stocksService         stocks.StocksService
}

// ProvideTransactionService is a method to handle DI
func ProvideTransactionService(r TransactionRepository, s stocks.StocksService) TransactionService {
	return TransactionService{r, s}
}

// GetAll returns all the transactions in the database
func (service *TransactionService) GetAll() []StockTransaction {
	return service.transactionRepository.getAll()
}

// AddTransaction adds a new set of transactions to the repositoryervice
func (service *TransactionService) AddTransaction(transaction TransactionDto) {
	stock := service.stocksService.Find(transaction.StockCode)

	transactions := transaction.Map()

	stockTransactions := make([]StockTransaction, len(transactions))
	for i, transaction := range transactions {
		stockTransaction := StockTransaction{}

		stockTransaction.BuyNew(&stock, &transaction)

		stockTransactions[i] = stockTransaction
	}

	service.transactionRepository.addTransactions(stockTransactions)
}
