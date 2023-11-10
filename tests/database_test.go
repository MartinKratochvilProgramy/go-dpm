package tests

import (
	"go-dpm/database"
	"testing"
)

func TestCreateUser(t *testing.T) {
	db := database.NewMemoryDatabase()

	err := db.SeedDatabase()
	if err != nil {
		t.Fatal("Failed to seed database: ", err)
	}

	err = db.CreateUser("Sbeve", "password", "CZK")
	if err != nil {
		t.Fatal("Failed to create user: ", err)
	}

	defer db.DB.Close()
}

func TestGetUser(t *testing.T) {
	db := database.NewMemoryDatabase()

	user, err := db.GetUser("Sbeve")
	t.Log("USER: ", user)
	if err != nil {
		t.Fatal("Failed to get user: ", err)
	}

	defer db.DB.Close()
}
