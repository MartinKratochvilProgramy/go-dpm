package database

import (
	"fmt"
	"os"  
	"database/sql"
)

type Database struct {
	DB *sql.DB
}

func NewDatabase() *Database {
	host := os.Getenv("database_host")
	port := os.Getenv("database_port")
	user := os.Getenv("database_user")
	password := os.Getenv("database_password")
	dbname := os.Getenv("database_dbname")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return &Database{DB: db}
}