package stocks

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// API is a set of methods for managing transactions
type API struct {
	service *Service
}

// ProvideStocksAPI provides a new instance for wire
func ProvideStocksAPI(s *Service) *API {
	return &API{service: s}
}

// GetAll returns all the stock objects in the database
func (a *API) GetAll(c *gin.Context) {
	stocks, err := a.service.GetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"stocks": stocks})
	return
}

// Find returns a single stock object with the provided code
func (a *API) Find(c *gin.Context) {
	var code string

	if err := c.ShouldBindJSON(&code); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	stocks, err := a.service.Find(code)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"stocks": stocks})
	return
}

// AddStock creates a new entry with the provided stock code
func (a *API) AddStock(c *gin.Context) {
	var code string

	if err := c.ShouldBindJSON(&code); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := a.service.AddStock(code)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
	return
}

// UploadStockHistory creates stock logs
func (a *API) UploadStockHistory(c *gin.Context) {
	var data StockLogListDto

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := a.service.AddStockLogs(data.Logs)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
	return
}
