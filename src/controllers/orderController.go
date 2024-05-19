package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/smtp"
	"shop/src/database"
	"shop/src/models"

	"github.com/gofiber/fiber/v3"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/checkout/session"
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

	tx := database.DB.Begin()

	if err := tx.Preload("User").First(&link).Error; err != nil {
		tx.Rollback()
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

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

	var lineItems []*stripe.CheckoutSessionLineItemParams

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

		if err := tx.Create(&item).Error; err != nil {
			tx.Rollback()
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		lineItems = append(lineItems, &stripe.CheckoutSessionLineItemParams{
			Name:        stripe.String(product.Title),
			Description: stripe.String(product.Description),
			Images:      []*string{stripe.String(product.Image)},
			Amount:      stripe.Int64(int64(product.Price * 100)),
			Currency:    stripe.String("usd"),
			Quantity:    stripe.Int64(int64(requestProduct["quantity"])),
		})
	}

	stripe.Key = "sk_test_51MsMZxHkyKSITrqJDufgIVwwj7cIBbFV4lO58MQfYPeVhhAskrrX8PkciAeUv9vPbLRaQsodsv1XjSm2X4u4zx5S00hwRfXjlP"

	params := stripe.CheckoutSessionParams{
		SuccessURL:         stripe.String("http://localhost:5000/success?source={CHECKOUT_SESSION_ID}"),
		CancelURL:          stripe.String("http://localhost:5000/error"),
		PaymentMethodTypes: stripe.StringSlice([]string{"card"}),
		LineItems:          lineItems,
	}

	source, err := session.New(&params)
	if err != nil {
		tx.Rollback()
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	order.TransactionId = source.ID

	if err := tx.Save(&order).Error; err != nil {
		tx.Rollback()
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	tx.Commit()

	return c.JSON(source)
}

type PlaceOrderData struct {
	Source string `json:"source"`
}

func PlaceOrder(c fiber.Ctx) error {
	var data PlaceOrderData
	err := json.Unmarshal(c.Body(), &data)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body format",
			"error":   err.Error(),
		})
	}

	order := models.Order{}
	database.DB.Preload("OrderItems").First(&order, models.Order{
		TransactionId: data.Source,
	})

	if order.ID == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "Order not found",
		})
	}

	order.Completed = true
	database.DB.Save(&order)

	go func(order models.Order) {
		ambassadorRevenue := 0.0
		adminRevenue := 0.0

		for _, orderItem := range order.OrderItems {
			ambassadorRevenue += orderItem.AmbassadorRevenue
			adminRevenue += orderItem.AdminRevenue
		}

		user := models.User{}
		user.ID = order.UserId

		database.DB.First(&user)

		database.Cache.ZIncrBy(context.Background(), "ambassadors_ranking", ambassadorRevenue, user.Name())

		ambassadorMessage := []byte(fmt.Sprintf("You have earned $%f revenue from the link #%s", ambassadorRevenue, order.Code))
		smtp.SendMail("host.docker.internal:1025", nil, "no-reply@ambassador.com", []string{order.AmbassadorEmail}, ambassadorMessage)

		adminMessage := []byte(fmt.Sprintf("Order #%d with a total of $%f has been completed", order.ID, adminRevenue))
		smtp.SendMail("host.docker.internal:1025", nil, "no-reply@ambassador.com", []string{"admin@localhost.com"}, adminMessage)

	}(order)

	return c.JSON(fiber.Map{
		"message": "Success",
	})
}
