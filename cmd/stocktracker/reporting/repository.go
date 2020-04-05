package reporting

import (
	"github.com/eddymoulton/stock-tracker/cmd/stocktracker/logger"
	"github.com/jinzhu/gorm"
)

// Repository is a set of methods for handling reporting database access
type Repository struct {
	db  *gorm.DB
	log *logger.Logger
}

// ProvideReportingRepository provides a new instance for wire
func ProvideReportingRepository(db *gorm.DB, logger *logger.Logger) *Repository {
	return &Repository{db, logger}
}
