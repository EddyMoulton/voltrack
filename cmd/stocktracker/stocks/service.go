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

// AddStockLogs create StockLogs for each passed data set
func (s *Service) AddStockLogs(data []StockLogDto) (int, error) {
	logs := make([]StockLog, len(data))

	existingStocks, err := s.GetAllCodes()

	if err != nil {
		s.log.Error((err.Error()))
		return 0, err
	}

	index := 0
	for _, newLog := range data {
		if (newLog.Date == time.Time{} || time.Now().Before(newLog.Date) || newLog.Value == 0) {
			s.log.Warning("Entry in uploaded log data is missing required information and will be skipped")
			s.log.Debug(fmt.Sprintf("Data: %+v", newLog))
		} else {
			stockCodeExists := false

			for _, code := range existingStocks {
				if newLog.StockCode == code {
					stockCodeExists = true
					break
				}
			}

			if stockCodeExists {
				logs[index] = stockLogFromDto(newLog)
				index++
			} else {
				s.log.Warning(fmt.Sprintf("Failed to add stock log for stock (%s) that is not in the system, skipping line", newLog.StockCode))
			}
		}
	}

	if index > 0 {
		logs = logs[0:index]
		return len(logs), s.repository.addStockLogs(logs)
	}

	s.log.Error("No valid stock logs to save to database")
	return 0, fmt.Errorf("No valid entries to save")
}

func (s *Service) GetLatestStockLog(stockCode string) (StockLog, error) {
	return s.repository.GetLatestStockLog(stockCode)
}
