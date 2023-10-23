package database

import "go-dpm/bcrypt"

func (d *Database) CreateUser(username string, password string, currency string) error {
	query := "INSERT INTO users (username, password_hash, currency) VALUES ($1, $2, $3);"

	passwordHash, err := bcrypt.HashPassword(password)
	if err != nil {
		return err
	}

	_, err = d.DB.Exec(query, username, passwordHash, currency)

	return err
}
