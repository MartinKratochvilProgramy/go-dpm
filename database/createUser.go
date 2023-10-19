package database

func (d *Database) CreateUser(username string, password string) error {
	query := "INSERT INTO users (username, password) VALUES ($1, $2);"

	_, err := d.DB.Exec(query, username, password)

	return err
}
