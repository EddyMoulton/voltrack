package stocks

// StockDTO is used when passing stock data to and from the API
type StockDTO struct {
	Code        string
	Description string
}

// Map converts a Stock instance to StockDTO
func (s *Stock) Map() (dto StockDTO) {
	dto = StockDTO{Code: s.Code, Description: s.Description}
	return dto
}
