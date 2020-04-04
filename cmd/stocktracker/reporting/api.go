package reporting

// API is a set of methods for accessing reporting
type API struct {
	service *Service
}

// ProvideReportingAPI provides a new instance for wire
func ProvideReportingAPI(s *Service) *API {
	return &API{service: s}
}
