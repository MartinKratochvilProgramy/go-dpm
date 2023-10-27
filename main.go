package main

import (
	"go-dpm/database"
	"go-dpm/router"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := database.NewDatabase()
	router := router.NewRouter(db)

	router.Run()
}
