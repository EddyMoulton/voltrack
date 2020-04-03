package stocks

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// API is a set of methods for managing transactions
type API struct {
	service Service
}

// ProvideStocksAPI provides a new instance for wire
func ProvideStocksAPI(s Service) API {
	return API{service: s}
}

// GetAll returns all the stock objects in the database
func (api *API) GetAll(c *gin.Context) {
	stocks, err := api.service.GetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"stocks": stocks})
	return
}

// Find returns a single stock object with the provided code
func (api *API) Find(c *gin.Context) {
	var code string

	if err := c.ShouldBindJSON(&code); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	stocks, err := api.service.Find(code)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"stocks": stocks})
	return
}

// AddStock creates a new entry with the provided stock code
func (api *API) AddStock(c *gin.Context) {
	var code string

	if err := c.ShouldBindJSON(&code); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := api.service.AddStock(code)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
	return
}
