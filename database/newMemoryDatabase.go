package database

import (
	"database/sql"

	_ "github.com/proullon/ramsql/driver"
)

func NewMemoryDatabase() *Database {
	db, err := sql.Open("ramsql", "TestLoadUserAddresses")

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return &Database{DB: db}
}
