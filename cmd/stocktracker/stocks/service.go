package stocks

import (
	"fmt"
	"time"

	"github.com/eddymoulton/stock-tracker/cmd/stocktracker/logger"
)

// Service is an object that provides methods for altering or manipulating stocks
type Service struct {
	repository *Repository
	exchanges  *Exchanges
	log        *logger.Logger
}

// ProvideStocksService is a method to handle DI
func ProvideStocksService(r *Repository, exchanges *Exchanges, logger *logger.Logger) *Service {
	return &Service{r, exchanges, logger}
}

// GetAll returns all the stock objects in the database
func (s *Service) GetAll() ([]Stock, error) {
	return s.repository.getAll()
}

// GetAllCodes returns all the stock codes in the database
func (s *Service) GetAllCodes() ([]string, error) {
	return s.repository.getAllCodes()
}

// Find returns a single stock object with the provided code
func (s *Service) Find(code string) (Stock, error) {
	return s.repository.find(code)
}

// AddStock creates a new entry with the provided stock code
func (s *Service) AddStock(code string) (Stock, error) {
	stock, err := s.exchanges.getStockPrice(code)

	if err != nil {
		s.log.Error("Could not get stock information for code", code, "cancelling add operation")
		return Stock{}, err
	}

	stockModel, err := s.repository.add(Stock{Code: code, Description: stock.Description})

	if err == nil {
		s.LogStocks([]string{stockModel.Code})
	}

	return stockModel, err
}

// LogAllStocks grabs the current price for all stocks in the database and creates StockLogs for each
func (s *Service) LogAllStocks() {
	stockCodes, err := s.GetAllCodes()

	if err == nil {
		s.LogStocks(stockCodes)
	}
}

// LogStocks grabs the current price for all passed stock codes and creates StockLogs for each
func (s *Service) LogStocks(stockCodes []string) {
	codes := make([]string, len(stockCodes))

	for i, code := range stockCodes {
		codes[i] = code
	}
	s.log.Info(fmt.Sprintf("Capturing value for stock codes: %v", codes))

	logs := make([]StockLog, len(codes))

	for i, code := range codes {
		result, err := s.exchanges.getStockPrice(code)

		if err != nil {
			s.log.Error(err.Error())
		}

		value := int64(result.LastPrice * 10000) // Convert to x10^4 int
		logs[i] = StockLog{Date: time.Now(), StockCode: code, Value: value}
	}

	s.repository.addStockLogs(logs)
}

func (s *Service) GetLatestStockLog(stockCode string) (StockLog, error) {
	return s.repository.GetLatestStockLog(stockCode)
}
