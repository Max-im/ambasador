package controllers

import (
	"shop/src/database"
	"shop/src/models"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

func Links(c fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var links []models.Link

	database.DB.Where("user_id = ?", id).Find(&links)

	for _, link := range links {
		var orders []models.Order
		database.DB.Where("code = ? and completed = true", link.Code).Find(&orders)
	}

	return c.JSON(links)
}
