package reporting

import (
	"net/http"

	"github.com/eddymoulton/stock-tracker/cmd/stocktracker/logger"
	"github.com/gin-gonic/gin"
)

// API is a set of methods for accessing reporting
type API struct {
	service *Service
	log     *logger.Logger
}

// ProvideReportingAPI provides a new instance for wire
func ProvideReportingAPI(s *Service, logger *logger.Logger) *API {
	return &API{s, logger}
}

// GetOwnedStockLogs returns all reporting logs for the provided codes in the date range
func (a *API) GetOwnedStockLogs(c *gin.Context) {
	var input StockRangeRequestDto

	if err := c.ShouldBindJSON(&input); err != nil {
		a.log.Warning(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	logs, err := a.service.GetOwnedStockLogs(input.StockCodes, input.Start, input.End)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"ownedStockLogs": logs})
	return
}

// GenerateSummaryLogs creates records for all stocks owned within the date range provided
func (a *API) GenerateSummaryLogs(c *gin.Context) {
	var dateRange DateRangeDTO

	if err := c.ShouldBindJSON(&dateRange); err != nil {
		a.log.Warning(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := a.service.GenerateSummaryLogs(dateRange.Start, dateRange.End)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
	return
}
