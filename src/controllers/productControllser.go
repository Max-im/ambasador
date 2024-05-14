package controllers

import (
	"encoding/json"
	"net/http"
	"shop/src/database"
	"shop/src/models"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

func Products(c fiber.Ctx) error {
	var products []models.Product

	database.DB.Find(&products)

	return c.JSON(products)
}

func CreateProduct(c fiber.Ctx) error {
	var product models.Product
	err := json.Unmarshal(c.Body(), &product)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body format",
			"error":   err.Error(),
		})
	}

	database.DB.Create(&product)

	return c.JSON(product)
}

func GetProduct(c fiber.Ctx) error {
	var product models.Product

	id, _ := strconv.Atoi(c.Params("id"))

	database.DB.Where("id = ?", id).First(&product)
	if product.ID == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "Product not found",
		})
	}
	return c.JSON(product)
}

func UpdateProduct(c fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	product := models.Product{
		ID: uint(id),
	}

	err := json.Unmarshal(c.Body(), &product)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body format",
			"error":   err.Error(),
		})
	}
	database.DB.Model(&product).Updates(&product)

	return c.JSON(product)
}

func DeleteProduct(c fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	database.DB.Where("id = ?", id).Delete(&models.Product{})

	return nil
}
