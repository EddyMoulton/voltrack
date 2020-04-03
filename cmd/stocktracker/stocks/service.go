package stocks

type StocksService struct {
	stocksRepository StocksRepository
}

func ProvideStocksService(r StocksRepository) StocksService {
	return StocksService{r}
}

// Stock CRUD
func (service *StocksService) GetAll() []Stock {
	return service.stocksRepository.getAll()
}

func (service *StocksService) Find(code string) Stock {
	return service.stocksRepository.find(code)
}

func (service *StocksService) AddStock(code string) {
	stock := GetStockPrice(code)

	service.stocksRepository.add(Stock{Code: code, Description: stock.Description})
}

// Stock Logging
func (service *StocksService) LogStocks() {
	stocks := service.GetAll()
	codes := make([]string, len(stocks))

	for i, stock := range stocks {
		codes[i] = stock.Code
	}

	logs := make([]StockLog, len(codes))

	for i, code := range codes {
		result := GetStockPrice(code)

		value := int64(result.LastPrice * 10000) // Convert to x10^4 int
		logs[i] = StockLog{StockCode: code, Value: value}
	}
}
