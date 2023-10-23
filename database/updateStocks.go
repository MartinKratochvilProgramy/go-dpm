package database

import (
	"fmt"
	"go-dpm/types"
	yfinanceapi "go-dpm/yFinanceAPI"
	"sync"
)

func (d *Database) UpdateStocks() error {
	stocks, err := d.GetStocks()
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	errCh := make(chan error, len(stocks))

	for _, stock := range stocks {
		wg.Add(1)
		go d.updateStock(stock, &wg, errCh)
	}

	wg.Wait()
	close(errCh)

	for err := range errCh {
		fmt.Println(err)
	}

	return nil
}

func (d *Database) updateStock(stock types.Stock, wg *sync.WaitGroup, errCh chan error) {
	defer wg.Done()

	quote, err := yfinanceapi.GetQuote(stock.Ticker)
	if err != nil {
		errCh <- fmt.Errorf("Error processing ticker %s %s", stock.Ticker, err)
	}

	stock.PrevClose = quote.PrevClose
	stock.CurrentPrice = quote.CurrentPrice

	_, err = d.DB.Exec("UPDATE stocks SET prev_close=$1 WHERE ticker=$2;", quote.PrevClose, stock.Ticker)
	if err != nil {
		errCh <- fmt.Errorf("Error processing ticker %s %s", stock.Ticker, err)
	}

	_, err = d.DB.Exec("UPDATE stocks SET current_price=$1 WHERE ticker=$2;", quote.CurrentPrice, stock.Ticker)
	if err != nil {
		errCh <- fmt.Errorf("Error processing ticker %s %s", stock.Ticker, err)
	}
}
