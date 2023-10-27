package database

import (
	"errors"
	"go-dpm/utils/bcrypt"
)

func (d *Database) CreateUser(username string, password string, currency string) error {
	// check if user is duplicate
	var duplicate = false
	err := d.DB.QueryRow("SELECT 1 FROM users WHERE username = $1 LIMIT 1", username).Scan(&duplicate)
	if duplicate {
		return errors.New("Username already exists.")
	}

	query := "INSERT INTO users (username, password_hash, currency) VALUES ($1, $2, $3);"

	passwordHash, err := bcrypt.HashPassword(password)
	if err != nil {
		return err
	}

	_, err = d.DB.Exec(query, username, passwordHash, currency)

	return err
}
