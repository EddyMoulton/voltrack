package transactions

import (
	"fmt"
	"sort"
	"time"

	"github.com/eddymoulton/voltrack/pkg/logger"
	"github.com/eddymoulton/voltrack/pkg/stocks"
)

// Service is an object that provides methods for altering or manipulating stock transactions
type Service struct {
	repository    *Repository
	stocksService *stocks.Service
	log           *logger.Logger
}

// ProvideTransactionsService is a method to handle DI
func ProvideTransactionsService(r *Repository, s *stocks.Service, logger *logger.Logger) *Service {
	return &Service{r, s, logger}
}

// GetAll returns all the transactions in the database
func (s *Service) GetAll() ([]StockTransaction, error) {
	return s.repository.getAll()
}

func (s *Service) getStockCodesInTransactions(transactions []StockTransaction) []string {
	stockCodes := make([]string, 0)
	for _, transaction := range transactions {
		addCode := true
		for _, code := range stockCodes {
			if transaction.StockCode == code {
				addCode = false
				break
			}
		}

		if addCode {
			stockCodes = append(stockCodes, transaction.StockCode)
		}
	}

	return stockCodes
}

// GetCurrentStocks returns all the stock objects in the database
func (s *Service) GetCurrentStocks() ([]OwnedStockSummaryDTO, error) {
	transactions, err := s.repository.getAllUnsoldStockTransactions()

	if err != nil {
		return make([]OwnedStockSummaryDTO, 0), err
	}

	stockCodes := s.getStockCodesInTransactions(transactions)

	stockLogs := make([]stocks.StockLog, 0, len(stockCodes))
	for _, code := range stockCodes {
		log, err := s.stocksService.GetLatestStockLog(code)

		if err != nil {
			return make([]OwnedStockSummaryDTO, 0), err
		}

		stockLogs = append(stockLogs, log)
	}

	result := createStockSummaries(stockCodes, transactions, stockLogs)

	return result, nil
}

// GetTransactionSummaries returns all the stock objects in the database
func (s *Service) GetTransactionSummaries() ([]TransactionSummaryDTO, error) {
	transactions, err := s.repository.getAllUnsoldStockTransactions()

	if err != nil {
		return make([]TransactionSummaryDTO, 0), err
	}

	stockCodes := s.getStockCodesInTransactions(transactions)

	summaries := make(map[string]map[int64]TransactionSummaryDTO)

	for _, code := range stockCodes {
		summaries[code] = make(map[int64]TransactionSummaryDTO)
	}

	latestStockLogs := make(map[string]int64)
	for _, code := range stockCodes {
		log, err := s.stocksService.GetLatestStockLog(code)

		if err != nil {
			return make([]TransactionSummaryDTO, 0), err
		}

		latestStockLogs[log.StockCode] = log.Value
	}

	for _, transaction := range transactions {
		temp := TransactionSummaryDTO{}
		if val, ok := summaries[transaction.StockCode][transaction.BuyTransactionID]; !ok {
			temp.Code = transaction.StockCode
			temp.Quantity = 1
			temp.Cost = transaction.BuyTransaction.Cost
			temp.Date = transaction.BuyTransaction.Date
			temp.Value = latestStockLogs[transaction.StockCode]
		} else {
			temp = val
			temp.Quantity++
		}

		summaries[transaction.StockCode][transaction.BuyTransactionID] = temp
	}

	result := make([]TransactionSummaryDTO, 0, len(summaries))

	for _, stock := range summaries {
		for _, value := range stock {
			result = append(result, value)
		}
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Date.Before(result[j].Date)
	})

	return result, nil
}

// AddBuyTransaction adds a new set of transactions to the repository
func (s *Service) AddBuyTransaction(transactionDTO TransactionDTO) {
	stock, err := s.stocksService.AddStock(transactionDTO.StockCode)

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

// AddBuyTransactions adds a new set of transactions to the repository
func (s *Service) AddBuyTransactions(data []TransactionDTO) (int, error) {

	allStockTransactions := []StockTransaction{}
	for _, transactionDto := range data {
		if (transactionDto.Date == time.Time{} || time.Now().Before(transactionDto.Date) || transactionDto.Cost == 0 || transactionDto.StockCode == "" || transactionDto.Quantity == 0) {
			s.log.Warning("Entry in uploaded log data is missing required information and will be skipped")
			s.log.Debug(fmt.Sprintf("Data: %+v", transactionDto))
		} else {
			stock, err := s.stocksService.AddStock(transactionDto.StockCode)

			if err == nil {
				transactions := transactionDto.Map()

				stockTransactions := make([]StockTransaction, len(transactions))
				for i, transaction := range transactions {
					stockTransaction := StockTransaction{}

					stockTransaction.BuyNew(&stock, &transaction)

					stockTransactions[i] = stockTransaction
				}

				allStockTransactions = append(allStockTransactions, stockTransactions...)
			}
		}
	}

	if len(allStockTransactions) > 0 {
		return len(allStockTransactions), s.repository.addTransactions(allStockTransactions)
	}

	s.log.Error("No valid stock logs to save to database")
	return 0, fmt.Errorf("No valid entries to save")
}

// AddSellTransaction adds a new set of transactions to the repository
func (s *Service) AddSellTransaction(transactionDTO TransactionDTO) {
	transactions := transactionDTO.Map()

	stockTransactions, err := s.repository.getOldestUnsoldStockTransactions(transactionDTO.StockCode, len(transactions))

	if err != nil {
		s.log.Error("Cannot find stock to add sale transaction for")
	}

	for i, transaction := range transactions {
		stockTransactions[i].Sell(&transaction)
	}

	s.repository.updateTransactions(stockTransactions)
}
