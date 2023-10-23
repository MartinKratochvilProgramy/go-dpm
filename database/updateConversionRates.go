package database

import yfinanceapi "go-dpm/yFinanceAPI"

func (d *Database) UpdateConversionRates() error {
	conversionRates, err := d.GetConversionRates()
	if err != nil {
		return err
	}

	for _, conversionRate := range conversionRates {
		rate, err := yfinanceapi.GetConversionRate(conversionRate.CurrencyPair)
		if err != nil {
			return err
		}

		_, err = d.DB.Exec("UPDATE conversion_rates SET conversion_rate=$1 WHERE currency_pair=$2;", rate, conversionRate.CurrencyPair)
		if err != nil {
			return err
		}

	}

	return nil
}
