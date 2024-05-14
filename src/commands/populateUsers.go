package main

import (
	"shop/src/database"
	"shop/src/models"

	"github.com/go-faker/faker/v4"
)

func main() {

	database.Connect()

	for i := 0; i < 30; i++ {
		ambassador := models.User{
			FirstName:    faker.FirstName(),
			LastName:     faker.LastName(),
			Email:        faker.Email(),
			IsAmbassador: true,
		}

		ambassador.SetPassword("123")
		database.DB.Create(&ambassador)
	}
}
