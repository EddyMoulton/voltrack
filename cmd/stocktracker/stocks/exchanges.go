package stocks

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://www.asx.com.au/asx/1/share/"

func GetStockPrice(stockCode string) AsxResult {
	resp, err := http.Get(url + stockCode + "/")

	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("ERROR: %v\n", err)
		}
		bodyString := string(bodyBytes)

		var result AsxResult
		json.Unmarshal([]byte(bodyString), &result)
		fmt.Println(result.LastPrice)

		return result
	}

	return AsxResult{}
}

type AsxResult struct {
	Code        string  `json:"code"`
	Description string  `json:"desc_full"`
	LastPrice   float64 `json:"last_price"`
}
