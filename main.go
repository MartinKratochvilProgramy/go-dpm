package main

import (
	"fmt"
	"go-dpm/bcrypt"
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

	user, err := db.GetUser("Martin")
	fmt.Println(err)

	fmt.Println(bcrypt.CheckPasswordHash("password", user.PasswordHash))

}
