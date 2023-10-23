package database

import (
	"log"

	"go-dpm/types"
)

func (d *Database) GetStockInPortfolio(username string, ticker string) (*types.StockInPortfolio, error) {
	query := `	
				SELECT
					s.ticker AS ticker,
					s.current_price AS current_price,
					sp.shares,
					s.current_price * sp.shares AS total
				FROM stocks_in_portfolio sp
				JOIN users u ON sp.user_id = u.id
				JOIN stocks s ON sp.stock_id = s.id
				WHERE u.username = $1 AND ticker = $2;
				`
	rows, err := d.DB.Query(query, username, ticker)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			ticker       string
			currentPrice float64
			shares       int
			total        float64
		)
		if err := rows.Scan(&ticker, &currentPrice, &shares, &total); err != nil {
			log.Fatal(err)
		}

		stockInPortfolio := types.StockInPortfolio{
			Ticker:       ticker,
			CurrentPrice: currentPrice,
			Shares:       shares,
			Total:        total,
		}

		return &stockInPortfolio, nil
	}

	return nil, nil
}
