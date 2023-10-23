package database

import (
	"go-dpm/types"

	"github.com/lib/pq"
)

func (d *Database) GetConversionRates() ([]types.ConversionRate, error) {
	rows, err := d.DB.Query("SELECT * FROM conversion_rates;")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var conversionRates []types.ConversionRate

	for rows.Next() {
		var (
			id           int
			currencyPair string
			rate         float64
			updatedAt    pq.NullTime
		)
		if err := rows.Scan(&id, &currencyPair, &rate, &updatedAt); err != nil {
			return nil, err
		}

		cr := types.ConversionRate{
			Id:           id,
			CurrencyPair: currencyPair,
			Rate:         rate,
		}

		conversionRates = append(conversionRates, cr)
	}

	return conversionRates, nil
}
