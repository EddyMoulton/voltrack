package stocks

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/eddymoulton/stock-tracker/cmd/stocktracker/logger"
)

// Exchanges is a set of methods for interacting with stock exchanges
type Exchanges struct {
	logger *logger.Logger
}

// ProvideExchanges provides a new instance for wire
func ProvideExchanges(logger *logger.Logger) Exchanges {
	return Exchanges{logger}
}

const url = "https://www.asx.com.au/asx/1/share/"

func getStockPrice(stockCode string) (AsxResult, error) {
	resp, err := http.Get(url + stockCode + "/")

	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return AsxResult{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("ERROR: %v\n", err)
			return AsxResult{}, err
		}
		bodyString := string(bodyBytes)

		var result AsxResult
		json.Unmarshal([]byte(bodyString), &result)
		fmt.Println(stockCode)
		fmt.Println(result.LastPrice)

		return result, nil
	}

	return AsxResult{}, fmt.Errorf("Failed to get AsxResults: %s", resp.Status)
}

// AsxResult is a object for passing returned data from Asx Stock Queries
type AsxResult struct {
	Code        string  `json:"code"`
	Description string  `json:"desc_full"`
	LastPrice   float64 `json:"last_price"`
}
