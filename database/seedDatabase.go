package database

func (d *Database) SeedDatabase() error {
	createTableSQL := `
		CREATE TABLE IF NOT EXISTS users(
			id BIGSERIAL PRIMARY KEY,
			username VARCHAR (100) UNIQUE NOT NULL,
			password_hash VARCHAR (100) NOT NULL,
			created_at TIMESTAMP DEFAULT NOW() NOT NULL,
			changed_at TIMESTAMP DEFAULT NULL,
			currency VARCHAR(10))
		`

	_, err := d.DB.Exec(createTableSQL)
	if err != nil {
		return err
	}

	return nil
}
