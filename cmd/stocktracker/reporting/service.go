package reporting

import (
	"strconv"
	"time"

	"github.com/eddymoulton/stock-tracker/cmd/stocktracker/helpers"
	"github.com/eddymoulton/stock-tracker/cmd/stocktracker/logger"
	"github.com/eddymoulton/stock-tracker/cmd/stocktracker/stocks"
	"github.com/eddymoulton/stock-tracker/cmd/stocktracker/transactions"
)

// Service is an object that provides methods for altering or manipulating reports
type Service struct {
	repository             *Repository
	transactionsRepository *transactions.Repository
	stocksRepository       *stocks.Repository
	log                    *logger.Logger
}

// ProvideReportingService is a method to handle DI
func ProvideReportingService(r *Repository,
	transactionsRepository *transactions.Repository,
	stocksRepository *stocks.Repository,
	logger *logger.Logger) *Service {

	return &Service{r, transactionsRepository, stocksRepository, logger}
}

// GenerateSummaryLogs takes all stock transactions that exist in the time period and create/replace
func (s *Service) GenerateSummaryLogs(start, end time.Time) error {
	transactions, err := s.transactionsRepository.GetStockTransactionsExistingBetween(start, end)

	if err == nil {
		s.log.Info("Creating", strconv.FormatInt(int64(len(transactions)), 10), "records")

		// Initialise map and date slice
		m := make(map[time.Time]map[string]int)
		dates := make([]time.Time, helpers.DaysBetweenDatesInclusive(start, end))
		codes := []string{}

		for rd := helpers.RangeDate(start, end); ; {
			date, index := rd()
			if date.IsZero() {
				break
			}

			m[date] = make(map[string]int)
			dates[index] = date
		}

		// Add counts of each stock per day of the range to the map
		for _, item := range transactions {
			for rd := helpers.RangeDate(start, end); ; {
				date, _ := rd()
				if date.IsZero() {
					break
				}

				if helpers.InTimeSpan(item.BuyTransaction.Date, item.SellTransaction.Date, date) {
					if _, ok := m[date][item.StockCode]; ok {
						m[date][item.StockCode]++
					} else {
						codes = append(codes, item.StockCode)
						m[date][item.StockCode] = 1
					}
				}
			}
		}

		// Get cost values for stocks over date range
		stockLogs, err := s.stocksRepository.GetStockLogs(codes, start, end)
		s.log.Debug("Number of stock logs", strconv.FormatInt(int64(len(stockLogs)), 10))

		if err != nil {
			return err
		}

		ownedStockLogs := make([]OwnedStockLog, len(codes)*len(dates))
		index := 0

		for date, stockMap := range m {
			for stockCode, quantity := range stockMap {
				var value int64

				for _, item := range stockLogs {
					if item.StockCode == stockCode && helpers.OnSameDay(item.Date, date) {
						value = item.Value
						break
					}
				}

				if date.IsZero() || stockCode == "" || quantity == 0 || value == 0 {
					s.log.Warning("Missing data for", stockCode, "on", date.Format("2006-01-02"))
				} else {
					ownedStockLogs[index] = OwnedStockLog{
						Date:            date,
						StockCode:       stockCode,
						Quantity:        int64(quantity),
						IndividualValue: value,
						TotalValue:      value * int64(quantity),
					}

					index++
				}
			}
		}

		ownedStockLogs = ownedStockLogs[0:index]

		s.log.Debug("SUMMARY")

		for _, log := range ownedStockLogs {
			s.log.Debug(log.StockCode, log.Date.Format("2006-01-02 15:04:05"), strconv.FormatInt(log.IndividualValue, 10), strconv.FormatInt(log.Quantity, 10), strconv.FormatInt(log.TotalValue, 10))
		}
	}

	return err
}

// TODO
// Take all stock transactions that exist in the time period, order by date (old -> new)
// Cycle through from start to end dates and generate count of each stock per day
// Correlate to price on each day for each stock
// Save back to database as:
// {
//	Stock
//	Date
// 	Price
// }

// Can then be charted with db queries for each
