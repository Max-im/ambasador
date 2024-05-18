package controllers

import (
	"shop/src/database"
	"shop/src/models"

	"github.com/gofiber/fiber/v3"
)

func Ambassadors(c fiber.Ctx) error {
	var users []models.User

	database.DB.Where("is_ambassador = true").Find(&users)

	return c.JSON(users)
}

func Ranking(c fiber.Ctx) error {
	var users []models.User

	database.DB.Where("is_ambassador = true").Find(&users)

	var result []interface{}

	for _, user := range users {
		ambassador := models.Ambassador(user)
		ambassador.CalculateRevenue(database.DB)
		result = append(result, fiber.Map{
			user.Name(): ambassador.Revenue,
		})
	}

	return c.JSON(result)
}
