package database

import (
	"errors"
	"go-dpm/types"

	"github.com/lib/pq"
)

func (d *Database) GetUser(username string) (*types.User, error) {
	rows, err := d.DB.Query(`SELECT * FROM users WHERE username = $1`, username)

	if err != nil {
		return nil, err
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
			return nil, err
		}

		user := types.User{
			Id:        id,
			Username:  username,
			Password:  password,
			CreatedAt: createdAt,
			ChangedAt: changedAt,
		}
		return &user, nil
	}

	return nil, errors.New("No user found!")
}
