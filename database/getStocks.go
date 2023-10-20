package database

import (
	"go-dpm/types"

	"github.com/lib/pq"
)

func (d *Database) GetStocks() ([]types.Stock, error) {
	rows, err := d.DB.Query("SELECT * FROM stocks")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stocks []types.Stock

	for rows.Next() {
		var (
			id           int
			ticker       string
			prevClose    float64
			currentPrice float64
			updatedAt    pq.NullTime
		)
		if err := rows.Scan(&id, &ticker, &prevClose, &currentPrice, &updatedAt); err != nil {
			return nil, err
		}

		stock := types.Stock{
			Id:           id,
			Ticker:       ticker,
			PrevClose:    prevClose,
			CurrentPrice: currentPrice,
			UpdatedAt:    updatedAt,
		}

		stocks = append(stocks, stock)
	}

	return stocks, nil
}
