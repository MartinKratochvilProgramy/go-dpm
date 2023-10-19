package main

import (
	"fmt"
	"go-orm/database"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := database.NewDatabase()

	user, err := db.GetUser("agfvztg")

	fmt.Println(user, err)
}
