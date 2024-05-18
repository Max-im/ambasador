package controllers

import (
	"encoding/json"
	"net/http"
	"shop/src/database"
	"shop/src/middlewares"
	"shop/src/models"
	"strconv"

	"github.com/go-faker/faker/v4"
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

type LinkRequestData struct {
	Products []int `json:"product"`
}

func CreateLink(c fiber.Ctx) error {
	var request LinkRequestData
	err := json.Unmarshal(c.Body(), &request)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body format",
			"error":   err.Error(),
		})
	}

	id, _ := middlewares.GetUserId(c)
	userId, _ := strconv.ParseUint(id, 10, 32)

	link := models.Link{
		Code:   faker.Username(),
		UserId: uint(userId),
	}

	for _, productId := range request.Products {
		product := models.Product{}
		product.ID = uint(productId)
		link.Products = append(link.Products, product)
	}

	database.DB.Create(&link)

	return c.JSON(link)
}

func Stats(c fiber.Ctx) error {
	id, _ := middlewares.GetUserId(c)
	userId, _ := strconv.ParseUint(id, 10, 32)

	var links []models.Link
	database.DB.Where("user_id = ?", userId).Find(&links)

	var result []interface{}

	var orders []models.Order

	for _, link := range links {
		database.DB.Preload("OrderItems").Where("code = ? and completed = true", link.Code).Find(&orders)
		revenue := 0.0

		for _, order := range orders {
			revenue += order.GetTotal()
		}

		result = append(result, fiber.Map{
			"code":    link.Code,
			"count":   len(orders),
			"revenue": revenue,
		})
	}
	return c.JSON(result)
}
