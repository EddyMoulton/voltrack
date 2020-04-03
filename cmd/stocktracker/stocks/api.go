package stocks

// API is a set of methods for managing transactions
type API struct {
	service Service
}

// ProvideStocksAPI provides a new instance for wire
func ProvideStocksAPI(s Service) API {
	return API{service: s}
}

// GetAll returns all the stock objects in the database
func (api *API) GetAll() ([]Stock, error) {
	return api.service.GetAll()
}

// Find returns a single stock object with the provided code
func (api *API) Find(code string) (Stock, error) {
	return api.service.Find(code)
}

// AddStock creates a new entry with the provided stock code
func (api *API) AddStock(code string) {
	api.service.AddStock(code)
}
