package database

import (
	yfinanceapi "go-dpm/yFinanceAPI"
)

func (d *Database) CreateStock(ticker string) error {
	quote, err := yfinanceapi.GetQuote(ticker)
	if err != nil {
		return err
	}

	query := "INSERT INTO stocks (ticker, prev_close, current_price, currency) VALUES ($1, $2, $3, $4);"

	_, err = d.DB.Exec(query, ticker, quote.PrevClose, quote.CurrentPrice, quote.Currency)

	return err
}
