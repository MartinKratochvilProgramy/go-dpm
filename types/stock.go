package types

import "github.com/lib/pq"

type StockInPortfolio struct {
	Ticker       string
	CurrentPrice float64
	Shares       int
	Total        float64
}

type Stock struct {
	Id           int
	Ticker       string
	PrevClose    float64
	CurrentPrice float64
	UpdatedAt    pq.NullTime
}
