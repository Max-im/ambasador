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
		var orderItems []models.OrderItem
		for j := 0; j < rand.Intn(5); j++ {
			price := float64(rand.Intn(90) + 10)
			qty := uint(rand.Intn(5) + 1)

			orderItems = append(orderItems, models.OrderItem{
				ProductTitle:      faker.Word(),
				Quantity:          qty,
				Price:             price,
				AdminRevenue:      price * 0.9 * float64(qty),
				AmbassadorRevenue: price * 0.1 * float64(qty),
				TotalPrice:        price,
			})

		}

		database.DB.Create(&models.Order{
			UserId:          uint(rand.Intn(30) + 1),
			Code:            faker.Username(),
			AmbassadorEmail: faker.Email(),
			FirstName:       faker.FirstName(),
			LastName:        faker.LastName(),
			Email:           faker.Email(),
			Completed:       true,
			OrderItem:       orderItems,
		})
	}
}
