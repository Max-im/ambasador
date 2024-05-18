package controllers

import (
	"encoding/json"
	"net/http"
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

type CreateOrderData struct {
	Code      string           `json:"code"`
	FirstName string           `json:"first_name"`
	LastName  string           `json:"last_name"`
	Email     string           `json:"email"`
	Address   string           `json:"address"`
	Country   string           `json:"country"`
	City      string           `json:"city"`
	Zip       string           `json:"zip"`
	Products  []map[string]int `json:"products"`
}

func CreateOrder(c fiber.Ctx) error {
	var data CreateOrderData
	err := json.Unmarshal(c.Body(), &data)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body format",
			"error":   err.Error(),
		})
	}

	link := models.Link{
		Code: data.Code,
	}

	database.DB.Preload("User").First(&link)
	if link.ID == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "Link not found",
		})
	}

	order := models.Order{
		Code:            link.Code,
		UserId:          link.UserId,
		AmbassadorEmail: link.User.Email,
		FirstName:       data.FirstName,
		LastName:        data.LastName,
		Email:           data.Email,
		Address:         data.Address,
		Country:         data.Country,
		City:            data.City,
		Zip:             data.Zip,
	}

	database.DB.Create(&order)

	for _, requestProduct := range data.Products {
		product := models.Product{}
		product.ID = uint(requestProduct["product_id"])
		database.DB.First(&product)
		total := product.Price * float64(requestProduct["quantity"])

		item := models.OrderItem{
			OrderId:           order.ID,
			ProductTitle:      product.Title,
			Price:             product.Price,
			Quantity:          uint(requestProduct["quantity"]),
			AmbassadorRevenue: 0.1 * total,
			AdminRevenue:      0.9 * total,
		}
		database.DB.Create(&item)
	}

	return c.JSON(order)
}
