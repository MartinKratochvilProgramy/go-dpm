package yfinanceapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func CheckCurrencyExists(currency string) error {
	url := fmt.Sprint("https://query1.finance.yahoo.com/v8/finance/chart/", currency, "=X")

	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	var responseData map[string]interface{}

	// Unmarshal the JSON data into the struct
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		return err
	}

	result := responseData["chart"].(map[string]interface{})["result"]
	if result == nil {
		return errors.New("Currency not found.")
	}

	return nil
}
