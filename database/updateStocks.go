package database

import yfinanceapi "go-dpm/yFinanceAPI"

func (d *Database) UpdateStocks() error {
	stocks, err := d.GetStocks()
	if err != nil {
		return err
	}

	for _, stock := range stocks {
		quote, err := yfinanceapi.GetQuote(stock.Ticker)
		if err != nil {
			return err
		}

		stock.PrevClose = quote.PrevClose
		stock.CurrentPrice = quote.CurrentPrice

		_, err = d.DB.Exec("UPDATE stocks SET prev_close=$1 WHERE ticker=$2;", quote.PrevClose, stock.Ticker)
		if err != nil {
			return err
		}

		_, err = d.DB.Exec("UPDATE stocks SET current_price=$1 WHERE ticker=$2;", quote.CurrentPrice, stock.Ticker)
		if err != nil {
			return err
		}
	}

	return nil
}
