package database

import (
	"fmt"
	"go-dpm/types"
	"sync"
)

func (d *Database) GetPortfolio(username string) ([]types.StockInPortfolio, error) {

	// get raw stocks without conversion rates
	stocksInPortfolio, err := d.getStocksInPortfolio(username)

	// get user currency
	user, err := d.GetUser(username)
	if err != nil {
		return nil, err
	}
	userCurrency := user.Currency

	// update stocks using current conversion rate
	var wg sync.WaitGroup
	errCh := make(chan error, len(stocksInPortfolio))

	for i := range stocksInPortfolio {
		wg.Add(1)
		go func(s *types.StockInPortfolio) {
			defer wg.Done()

			if userCurrency != s.Currency {
				conversionRate, err := d.GetConversionRate(s.Currency, userCurrency)
				if err != nil {
					errCh <- fmt.Errorf("Error updating stock while getting portfolio %s %s", s.Ticker, err)
				}

				s.CurrentPrice = s.CurrentPrice * conversionRate
			}
		}(&stocksInPortfolio[i])
	}

	wg.Wait()
	close(errCh)

	for err := range errCh {
		fmt.Println(err)
	}

	return stocksInPortfolio, nil
}

func (d *Database) getStocksInPortfolio(username string) ([]types.StockInPortfolio, error) {
	// fetch raw stocks from db, stocks have different currencies
	query := `	
	SELECT
		s.ticker AS stock_ticker,
		s.current_price AS current_price,
		sp.shares,
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
			currency     string
		)
		if err := rows.Scan(&ticker, &currentPrice, &shares, &currency); err != nil {
			fmt.Println(err)
			return nil, err
		}

		stock := types.StockInPortfolio{
			Ticker:       ticker,
			CurrentPrice: currentPrice,
			Shares:       shares,
			Currency:     currency,
		}

		stocksInPortfolio = append(stocksInPortfolio, stock)
	}

	return stocksInPortfolio, nil
}
