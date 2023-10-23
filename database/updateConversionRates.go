package database

import (
	"fmt"
	"go-dpm/types"
	yfinanceapi "go-dpm/yFinanceAPI"
	"sync"
)

func (d *Database) UpdateConversionRates() error {
	conversionRates, err := d.GetConversionRates()
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	errCh := make(chan error, len(conversionRates))

	for _, conversionRate := range conversionRates {
		wg.Add(1)
		go d.updateConversionRate(conversionRate, &wg, errCh)
	}

	wg.Wait()
	close(errCh)

	for err := range errCh {
		fmt.Println(err)
	}

	return nil
}

func (d *Database) updateConversionRate(conversionRate types.ConversionRate, wg *sync.WaitGroup, errCh chan error) {
	defer wg.Done()

	rate, err := yfinanceapi.GetConversionRate(conversionRate.CurrencyPair)
	if err != nil {
		errCh <- fmt.Errorf("Error updating conversion rate ticker %s %s", conversionRate.CurrencyPair, err)
	}

	_, err = d.DB.Exec("UPDATE conversion_rates SET conversion_rate=$1 WHERE currency_pair=$2;", rate, conversionRate.CurrencyPair)
	if err != nil {
		errCh <- fmt.Errorf("Error updating conversion rate ticker %s %s", conversionRate.CurrencyPair, err)
	}
}
