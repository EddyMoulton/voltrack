package stocks

import (
	"fmt"

	"github.com/eddymoulton/stock-tracker/cmd/stocktracker/logger"
)

// Service is an object that provides methods for altering or manipulating stocks
type Service struct {
	stocksRepository *Repository
	exchanges        *Exchanges
	logger           *logger.Logger
}

// ProvideStocksService is a method to handle DI
func ProvideStocksService(r *Repository, exchanges *Exchanges, logger *logger.Logger) *Service {
	return &Service{r, exchanges, logger}
}

// GetAll returns all the stock objects in the database
func (service *Service) GetAll() ([]Stock, error) {
	return service.stocksRepository.getAll()
}

// Find returns a single stock object with the provided code
func (service *Service) Find(code string) (Stock, error) {
	return service.stocksRepository.find(code)
}

// AddStock creates a new entry with the provided stock code
func (service *Service) AddStock(code string) error {
	stock, err := service.exchanges.getStockPrice(code)

	if err != nil {
		return err
	}

	service.stocksRepository.add(Stock{Code: code, Description: stock.Description})

	return nil
}

// LogStocks grabs the current price for all stocks in the database and creates StockLogs for each
func (service *Service) LogStocks() {
	stocks, err := service.GetAll()

	if err == nil {
		codes := make([]string, len(stocks))

		for i, stock := range stocks {
			codes[i] = stock.Code
		}
		service.logger.LogInfo(fmt.Sprintf("Capturing value for stock codes: %v", codes))

		logs := make([]StockLog, len(codes))

		for i, code := range codes {
			result, err := service.exchanges.getStockPrice(code)

			if err != nil {
				service.logger.LogFatal(err.Error())
			}

			value := int64(result.LastPrice * 10000) // Convert to x10^4 int
			logs[i] = StockLog{StockCode: code, Value: value}
		}

		service.stocksRepository.addStockLogs(logs)
	}
}
