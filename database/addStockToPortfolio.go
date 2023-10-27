package database

import "fmt"

func (d *Database) AddStockToPortfolio(username string, ticker string, shares int) error {
	stock, err := d.GetStock(ticker)
	if err != nil {
		return err
	}
	if stock == nil {
		// create new stock write
		err = d.CreateStock(ticker)
		if err != nil {
			return err
		}
		fmt.Println("Created new stock " + ticker)
	}

	stockInPortfolio, err := d.GetStockInPortfolio(username, ticker)
	if err != nil {
		return err
	}

	if stockInPortfolio == nil {
		// create new portfolio write
		query := `
		WITH user_stock_ids AS (
			SELECT
			(SELECT id FROM users WHERE username = $1) AS user_id,
			(SELECT id FROM stocks WHERE ticker = $2) AS stock_id
		)
			INSERT INTO stocks_in_portfolio (user_id, stock_id, shares)
			VALUES (
				(SELECT user_id FROM user_stock_ids), 
				(SELECT stock_id FROM user_stock_ids), 
				$3
				);`

		_, err = d.DB.Exec(query, username, ticker, shares)
		return err

	} else {
		// add to existing stock write
		newAmount := stockInPortfolio.Shares + shares
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
	}
}
