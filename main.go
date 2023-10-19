package main

import (
	"fmt"
	"go-orm/database"
	"log"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := database.NewDatabase()

	stocks, err := db.GetPortfolio("Sbeve")

	fmt.Println(stocks)
}
