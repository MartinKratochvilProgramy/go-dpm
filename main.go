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

	// err = db.AddStockToPortfolio("Sbeve", "TABAK.PR", 5)
	// fmt.Println(err)

	pf, err := db.GetPortfolio("Sbeve")
	fmt.Println(pf)
}
