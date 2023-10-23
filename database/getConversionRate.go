package database

import (
	"fmt"
	yfinanceapi "go-dpm/yFinanceAPI"

	"github.com/lib/pq"
)

func (d *Database) GetConversionRate(originalCurrency string, targetCurrency string) (float64, error) {
	// store conversion rates in db so the would not have to be refetched every time

	currencyPair := fmt.Sprint(originalCurrency, targetCurrency)

	rows, err := d.DB.Query(`SELECT * FROM conversion_rates WHERE currency_pair = $1;`, currencyPair)
	if err != nil {
		return 0, err
	}

	for rows.Next() {
		var (
			id             int
			currencyPair   string
			conversionRate float64
			updatedAt      pq.NullTime
		)

		if err := rows.Scan(&id, &currencyPair, &conversionRate, &updatedAt); err != nil {
			return 0, err
		}

		return conversionRate, nil
	}

	// no conversion rate found in database, create new
	conversionRate, err := yfinanceapi.GetConversionRate(fmt.Sprint(originalCurrency, targetCurrency))
	if err != nil {
		return 0, err
	}

	err = d.createNewConversionRate(currencyPair, conversionRate)
	if err != nil {
		return 0, err
	}

	return conversionRate, nil
}

func (d *Database) createNewConversionRate(currencyPair string, conversionRate float64) error {
	query := "INSERT INTO conversion_rates (currency_pair, conversion_rate) VALUES ($1, $2);"

	_, err := d.DB.Exec(query, currencyPair, conversionRate)

	return err
}
