package database

import (
	"errors"
	"go-dpm/types"
)

func (d *Database) GetUser(username string) (*types.User, error) {
	rows, err := d.DB.Query(`SELECT username, password_hash, currency FROM users WHERE username = $1`, username)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			username string
			password string
			currency string
		)
		if err := rows.Scan(&username, &password, &currency); err != nil {
			return nil, err
		}

		user := types.User{
			Username:     username,
			PasswordHash: password,
			Currency:     currency,
		}
		return &user, nil
	}

	return nil, errors.New("No user found!")
}
