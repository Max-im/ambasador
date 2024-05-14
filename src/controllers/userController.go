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
