package main

import (
	"fmt"
	"go-dpm/database"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := database.NewDatabase()

	portfolio, err := db.GetPortfolio("Sbeve")
	fmt.Println(portfolio)
}
