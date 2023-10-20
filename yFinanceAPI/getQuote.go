package yfinanceapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Quote struct {
	PrevClose    float64
	CurrentPrice float64
}

func GetQuote(ticker string) (*Quote, error) {
	url := fmt.Sprint("https://query1.finance.yahoo.com/v8/finance/chart/", ticker)

	// Send an HTTP GET request
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var responseData map[string]interface{}

	// Unmarshal the JSON data into the struct
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		return nil, err
	}

	prevClose := responseData["chart"].(map[string]interface{})["result"].([]interface{})[0].(map[string]interface{})["meta"].(map[string]interface{})["chartPreviousClose"]
	currentPrice := responseData["chart"].(map[string]interface{})["result"].([]interface{})[0].(map[string]interface{})["meta"].(map[string]interface{})["regularMarketPrice"]

	quote := Quote{
		PrevClose:    prevClose.(float64),
		CurrentPrice: currentPrice.(float64),
	}

	return &quote, nil
}
