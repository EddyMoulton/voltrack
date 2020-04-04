package reporting

import (
	"github.com/eddymoulton/stock-tracker/cmd/stocktracker/logger"
)

// Service is an object that provides methods for altering or manipulating reports
type Service struct {
	repository *Repository
	logger     *logger.Logger
}

// ProvideReportingService is a method to handle DI
func ProvideReportingService(r *Repository, logger *logger.Logger) *Service {
	return &Service{r, logger}
}
