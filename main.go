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

	db.UpdateStocks()

	ts := time.Now()
	pf, err := db.GetPortfolio("Sbeve")

	elapsed := time.Since(ts)
	fmt.Println(err, elapsed)
	fmt.Println(pf)
}
