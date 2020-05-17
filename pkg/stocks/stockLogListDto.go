package stocks

// StockLogListDto is for transfering a list of StockLogDto
type StockLogListDto struct {
	Logs []StockLogDto `json:"logs" binding:"required"` // List of logs
}
