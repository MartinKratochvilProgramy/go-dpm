package yfinanceapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetConversionRate(currencyPair string) (float64, error) {
	url := fmt.Sprint("https://query1.finance.yahoo.com/v8/finance/chart/", currencyPair, "=X")

	response, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return 0, err
	}

	var responseData map[string]interface{}

	// Unmarshal the JSON data into the struct
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		return 0, err
	}

	prevClose := responseData["chart"].(map[string]interface{})["result"].([]interface{})[0].(map[string]interface{})["meta"].(map[string]interface{})["chartPreviousClose"]

	return prevClose.(float64), nil
}
