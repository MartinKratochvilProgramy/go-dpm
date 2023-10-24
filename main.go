package main

import (
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
	// router := router.NewRouter()

	err = db.RemoveUnusedStocks()

	// router.Run()
}
