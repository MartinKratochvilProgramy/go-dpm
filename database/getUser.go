package database

import (
	"errors"
	"go-dpm/types"
)

func (d *Database) GetUser(username string) (*types.User, error) {
	rows, err := d.DB.Query(`SELECT id, username, password_hash, currency FROM users WHERE username = $1`, username)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			id       int
			username string
			password string
			currency string
		)
		if err := rows.Scan(&id, &username, &password, &currency); err != nil {
			return nil, err
		}

		user := types.User{
			Id:           id,
			Username:     username,
			PasswordHash: password,
			Currency:     currency,
		}
		return &user, nil
	}

	return nil, errors.New("User not found.")
}
