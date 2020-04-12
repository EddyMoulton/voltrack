package transactions

import "github.com/eddymoulton/stock-tracker/cmd/stocktracker/stocks"

// OwnedStockSummaryDTO is used when returning summary data about an owned stock
type OwnedStockSummaryDTO struct {
	Code         string `json:"code" binding:"required"`
	Quanity      int    `json:"quanity" binding:"required"`
	CurrentValue int64  `json:"currentValue" binding:"required"`
	TotalValue   int64  `json:"totalValue" binding:"required"`
	PaidValue    int64  `json:"paidValue" binding:"required"`
	Difference   int64  `json:"difference" binding:"required"`
}

func CreateStockSummaries(stockCodes []string, stockTransactions []StockTransaction, latestPrices []stocks.StockLog) []OwnedStockSummaryDTO {
	summaries := make(map[string]OwnedStockSummaryDTO)

	for _, code := range stockCodes {
		summaries[code] = OwnedStockSummaryDTO{Code: code}
	}

	for _, transaction := range stockTransactions {
		temp := summaries[transaction.StockCode]
		temp.Quanity++
		temp.PaidValue += transaction.BuyTransaction.Cost
		summaries[transaction.StockCode] = temp
	}

	for _, latest := range latestPrices {
		temp := summaries[latest.StockCode]
		temp.CurrentValue = latest.Value
		temp.TotalValue = int64(temp.Quanity) * temp.CurrentValue
		temp.Difference = temp.TotalValue - temp.PaidValue
		summaries[latest.StockCode] = temp
	}

	result := make([]OwnedStockSummaryDTO, 0, len(summaries))

	for _, value := range summaries {
		result = append(result, value)
	}

	return result
}
