package database

import (
	"go-dpm/types"
)

func (d *Database) GetPortfolio(username string) ([]types.StockInPortfolio, error) {
	query := `	
		SELECT
			s.ticker AS stock_ticker,
			s.current_price AS current_price,
			sp.shares,
			s.current_price * sp.shares AS total,
			s.currency AS currency
		FROM
			stocks_in_portfolio sp
		JOIN
			users u ON sp.user_id = u.id
		JOIN
			stocks s ON sp.stock_id = s.id
		WHERE
			u.username = $1;
		`
	rows, err := d.DB.Query(query, username)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stocksInPortfolio []types.StockInPortfolio

	for rows.Next() {
		var (
			ticker       string
			currentPrice float64
			shares       int
			total        float64
			currency     string
		)
		if err := rows.Scan(&ticker, &currentPrice, &shares, &total, &currency); err != nil {
			return nil, err
		}

		stock := types.StockInPortfolio{
			Ticker:       ticker,
			CurrentPrice: currentPrice,
			Shares:       shares,
			Total:        total,
			Currency:     currency,
		}

		stocksInPortfolio = append(stocksInPortfolio, stock)
	}

	user, err := d.GetUser(username)
	if err != nil {
		return nil, err
	}
	userCurrency := user.Currency

	for i := range stocksInPortfolio {
		if userCurrency != stocksInPortfolio[i].Currency {
			conversionRate, err := d.GetConversionRate(stocksInPortfolio[i].Currency, userCurrency)
			if err != nil {
				return nil, err
			}

			stocksInPortfolio[i].CurrentPrice *= conversionRate
			stocksInPortfolio[i].Total *= conversionRate
		}
	}

	return stocksInPortfolio, nil
}
