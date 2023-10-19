package database

import (
	"fmt"
	"log"

	"github.com/lib/pq"
)

func (d *Database) GetUsers() {
	rows, err := d.DB.Query("SELECT * FROM users")

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			id         int
			username   string
			password   string
			created_at pq.NullTime
			changed_at pq.NullTime
		)
		if err := rows.Scan(&id, &username, &password, &created_at, &changed_at); err != nil {
			log.Fatal(err)
		}

		// Process the data from the row
		fmt.Printf("%d %s\n", id, username)
	}
}
