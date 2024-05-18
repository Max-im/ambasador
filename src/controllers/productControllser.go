package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"shop/src/database"
	"shop/src/models"
	"sort"
	"strconv"
	"strings"
	"time"

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
	go database.ClearCache("products_frontend", "products_backend")

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
	product := models.Product{}
	product.ID = uint(id)

	err := json.Unmarshal(c.Body(), &product)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body format",
			"error":   err.Error(),
		})
	}
	database.DB.Model(&product).Updates(&product)
	go database.ClearCache("products_frontend", "products_backend")
	return c.JSON(product)
}

func DeleteProduct(c fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	database.DB.Where("id = ?", id).Delete(&models.Product{})
	go database.ClearCache("products_frontend", "products_backend")

	return nil
}

func ProductsFrontend(c fiber.Ctx) error {
	var products []models.Product

	ctx := context.Background()

	result, err := database.Cache.Get(ctx, "products_frontend").Result()

	if err != nil {
		database.DB.Find(&products)
		bytes, _ := json.Marshal(products)

		database.Cache.Set(ctx, "products_frontend", bytes, 30*time.Minute)
	} else {
		json.Unmarshal([]byte(result), &products)
	}

	return c.JSON(products)
}

func ProductsBackend(c fiber.Ctx) error {
	var products []models.Product

	ctx := context.Background()

	result, err := database.Cache.Get(ctx, "products_backend").Result()

	if err != nil {
		database.DB.Find(&products)
		bytes, _ := json.Marshal(products)

		database.Cache.Set(ctx, "products_backend", bytes, 30*time.Minute)
	} else {
		json.Unmarshal([]byte(result), &products)
	}

	var seachResults []models.Product

	if s := c.Query("s"); s != "" {
		for _, product := range products {
			if strings.Contains(strings.ToLower(product.Title), strings.ToLower(s)) ||
				strings.Contains(strings.ToLower(product.Description), strings.ToLower(s)) {
				seachResults = append(seachResults, product)
			}
		}
		products = seachResults
	}

	if sortParam := c.Query("sort"); sortParam != "" {
		sortLower := strings.ToLower(sortParam)
		if sortLower == "asc" {
			sort.Slice(products, func(i, j int) bool {
				return products[i].Price < products[j].Price
			})
		} else if sortLower == "desc" {
			sort.Slice(products, func(i, j int) bool {
				return products[i].Price > products[j].Price
			})
		}
	}

	var total = len(products)
	page, _ := strconv.Atoi(c.Query("page", "1"))
	productsPerPage := 9

	var data []models.Product

	if total <= page*productsPerPage && total >= (page-1)*productsPerPage {
		data = products[(page-1)*productsPerPage : total]
	} else if total >= page*productsPerPage {
		data = products[(page-1)*productsPerPage : page*productsPerPage]
	} else {
		data = []models.Product{}
	}

	return c.JSON(fiber.Map{
		"data":      data,
		"total":     total,
		"page":      page,
		"last_page": total/productsPerPage + 1,
	})
}
