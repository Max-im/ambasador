package controllers

import (
	"shop/src/database"
	"shop/src/models"

	"github.com/gofiber/fiber/v3"
)

func Orders(c fiber.Ctx) error {
	var orders []models.Order

	database.DB.Preload("OrderItems").Find(&orders)

	for i, order := range orders {
		orders[i].Name = order.GetName()
		orders[i].Total = order.GetTotal()
	}

	return c.JSON(orders)
}
