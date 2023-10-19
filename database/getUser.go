package database

import (
	"go-orm/types"
	"log"

	"github.com/lib/pq"
)

func (d *Database) GetUser(username string) *types.User {
	rows, err := d.DB.Query(`SELECT * FROM users WHERE username = $1`, username)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			id        int
			username  string
			password  string
			changedAt pq.NullTime
			createdAt pq.NullTime
		)
		if err := rows.Scan(&id, &username, &password, &changedAt, &createdAt); err != nil {
			log.Fatal(err)
		}

		user := types.User{
			Id:        id,
			Username:  username,
			Password:  password,
			CreatedAt: createdAt,
			ChangedAt: changedAt,
		}
		return &user
	}

	return nil
}
