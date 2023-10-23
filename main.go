package main

import (
	"fmt"
	"go-dpm/database"
	"log"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := database.NewDatabase()

	ts := time.Now()
	err = db.UpdateStocks()

	elapsed := time.Since(ts)
	fmt.Println(err, elapsed)
}
