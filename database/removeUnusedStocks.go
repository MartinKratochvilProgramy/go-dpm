package database

import "fmt"

func (d *Database) RemoveUnusedStocks() error {
	// loop through stocks, if stock is not referenced in any portfolio, remove it

	stocks, err := d.GetStocks()
	if err != nil {
		return err
	}

	for _, stock := range stocks {
		rows, err := d.DB.Query("SELECT COUNT(*) FROM stocks_in_portfolio WHERE stock_id = $1;", stock.Id)
		if err != nil {
			return err
		}

		for rows.Next() {
			// get count
			var count int

			if err := rows.Scan(&count); err != nil {
				return err
			}

			if count == 0 {
				// if stock is unused, delete it
				err = d.removeStock(stock.Id)
				if err != nil {
					return err
				}
				fmt.Println("Removed " + stock.Ticker)
			}
		}
	}

	return nil
}

func (d *Database) removeStock(stockId int) error {
	query := `DELETE FROM stocks WHERE id = $1`
	_, err := d.DB.Exec(query, stockId)

	return err
}
