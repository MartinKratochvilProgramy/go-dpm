package database

import (
	"errors"
	"go-dpm/types"

	"github.com/lib/pq"
)

func (d *Database) GetStock(ticker string) (*types.Stock, error) {
	rows, err := d.DB.Query("SELECT * FROM stocks WHERE ticker=$1", ticker)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

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
		return &stock, nil
	}

	return nil, errors.New(ticker + " not found in stocks table!")
}
