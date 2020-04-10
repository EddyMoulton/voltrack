package reporting

import (
	"fmt"
	"strconv"
	"time"

	"github.com/eddymoulton/stock-tracker/cmd/stocktracker/helpers"
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

// GetOwnedStockLogs returns all reporting logs for the provided codes in the date range
func (r *Repository) GetOwnedStockLogs(stockCodes []string, start, end time.Time) ([]OwnedStockLog, error) {
	ownedStockLogs := []OwnedStockLog{}
	if err := r.db.
		Where("stock_code IN (?)", stockCodes).
		Where("date >= ?", helpers.RemoveTime(start)).
		Where("date <= ?", helpers.RemoveTime(end).Add(24*time.Hour)).
		Find(&ownedStockLogs).Error; err != nil {

		r.log.Warning(err.Error())
		return ownedStockLogs, err
	}

	return ownedStockLogs, nil
}

// SaveOwnedStockLogs adds (or updates existing) reporting logs
func (r *Repository) SaveOwnedStockLogs(logs []OwnedStockLog) error {
	r.log.DbAccess(fmt.Sprintf("Adding/Updating OwnedStockLogs (%d entries)", len(logs)))

	// Pull information about
	stockCodes := []string{}
	start := helpers.MaxTime()
	end := time.Time{}

	for _, log := range logs {
		addCode := true

		for _, code := range stockCodes {
			if log.StockCode == code {
				addCode = false
				break
			}
		}

		if addCode {
			stockCodes = append(stockCodes, log.StockCode)
		}

		if log.Date.Before(start) {
			start = log.Date
		}

		if log.Date.After(end) {
			end = log.Date
		}
	}

	// Load existing logs within date range
	existingOwnedStockLogs, err := r.GetOwnedStockLogs(stockCodes, start, end)

	if err != nil {
		return err
	}

	// Separate logs into existing and new logs
	newOwnedStockLogs := make([]OwnedStockLog, len(logs))
	newLogCount := 0
	for _, log := range logs {
		isNew := true

		for _, existingLog := range existingOwnedStockLogs {
			if existingLog.StockCode == log.StockCode && helpers.RemoveTime(existingLog.Date) == helpers.RemoveTime(log.Date) {
				isNew = false

				existingLog.update(log.Quantity, log.IndividualValue)

				break
			}
		}

		if isNew {
			r.log.Debug("Found new for", log.StockCode, log.Date.Format("2006-01-02"))
			newOwnedStockLogs[newLogCount] = log
			newLogCount++
		}
	}

	newOwnedStockLogs = newOwnedStockLogs[0:newLogCount]

	tx := r.db.Begin()

	r.log.Debug(strconv.FormatInt(int64(len(newOwnedStockLogs)), 10), "new entries")
	for _, log := range newOwnedStockLogs {
		if err := tx.Create(&log).Error; err != nil {
			r.log.Error(err.Error())
			r.log.DbAccess(fmt.Sprintf("Failed adding log for %s on %s, rolling back", log.StockCode, log.Date.Format("2006-01-02")))
			tx.Rollback()
			return err
		}
	}

	r.log.Debug(strconv.FormatInt(int64(len(existingOwnedStockLogs)), 10), "updated entries")
	for _, log := range existingOwnedStockLogs {
		if err := tx.Save(&log).Error; err != nil {
			r.log.Error(err.Error())
			r.log.DbAccess(fmt.Sprintf("Failed adding log for %s on %s, rolling back", log.StockCode, log.Date.Format("2006-01-02")))
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}
