package database

import (
	"errors"
	"fmt"
)

func (d *Database) RemoveStockFromPortfolio(username string, ticker string, sharesToSubtract int) error {
	stockInPortfolio, err := d.GetStockInPortfolio(username, ticker)
	if err != nil {
		return err
	}
	if stockInPortfolio == nil {
		errMessage := fmt.Sprintf("Ticker %s not found in %s's portfoio.", ticker, username)
		return errors.New(errMessage)
	}

	newAmount := stockInPortfolio.Shares - sharesToSubtract

	if newAmount > 0 {
		// set new stock shares in portfolio
		query := `
			WITH user_stock_ids AS (
				SELECT
				(SELECT id FROM users WHERE username = $1) AS user_id,
				(SELECT id FROM stocks WHERE ticker = $2) AS stock_id
			)
			UPDATE stocks_in_portfolio 
			SET shares = $3
			WHERE 
				user_id = (SELECT user_id FROM user_stock_ids)
				AND
				stock_id = (SELECT stock_id FROM user_stock_ids) 
			;`

		_, err = d.DB.Exec(query, username, ticker, newAmount)
		return err

	} else {
		// remove stock from portfolio
		query := `
			WITH user_stock_ids AS (
				SELECT
				(SELECT id FROM users WHERE username = $1) AS user_id,
				(SELECT id FROM stocks WHERE ticker = $2) AS stock_id
			)
			DELETE FROM stocks_in_portfolio 
			WHERE 
				user_id = (SELECT user_id FROM user_stock_ids)
				AND
				stock_id = (SELECT stock_id FROM user_stock_ids) 
			;`

		_, err = d.DB.Exec(query, username, ticker)
		return err

	}
}
