package main

import (
  "database/sql"
  "fmt"
  "log"

  _ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5433
	user     = "postgres"
	password = "password"
	dbname   = "mydb"
  )

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	  "password=%s dbname=%s sslmode=disable",
	  host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
	  panic(err)
	}
	defer db.Close()
  
	err = db.Ping()
	if err != nil {
	  panic(err)
	}

	// Perform a SELECT query
	rows, err := db.Query("SELECT * FROM users WHERE first_name='Sjoe'")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
        var (
            id int
            first_name string
			created_at string
            // Define variables for each column's data type
        )
        if err := rows.Scan(&id, &first_name, &created_at); err != nil {
            log.Fatal(err)
        }

        // Process the data from the row
        fmt.Printf("%d %s %s\n", id, first_name, created_at)
    }
}