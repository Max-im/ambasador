package main

import (
	"math/rand"
	"shop/src/database"
	"shop/src/models"

	"github.com/go-faker/faker/v4"
)

func main() {

	database.Connect()

	for i := 0; i < 30; i++ {
		product := models.Product{
			Title:       faker.Username(),
			Description: faker.Username(),
			Image:       faker.URL(),
			Price:       rand.Float64(),
		}

		database.DB.Create(&product)
	}
}
