package controllers

import (
	"encoding/json"
	"net/http"
	"shop/src/database"
	"shop/src/models"

	"github.com/gofiber/fiber/v3"
	"golang.org/x/crypto/bcrypt"
)

type RegisterRequestData struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirm"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
}

func Register(c fiber.Ctx) error {

	var data RegisterRequestData
	err := json.Unmarshal(c.Body(), &data)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body format",
			"error":   err.Error(),
		})
	}

	if data.Email == "" || data.Password == "" || data.PasswordConfirm == "" || data.FirstName == "" || data.LastName == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Missing required fields: 'email', 'password', 'first_name', 'last_name', or 'password confirm'",
		})
	}

	if data.Password != data.PasswordConfirm {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Passwords do not match",
		})
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data.Password), 12)

	user := models.User{
		Email:        data.Email,
		FirstName:    data.FirstName,
		LastName:     data.LastName,
		Password:     password,
		IsAmbassador: false,
	}

	database.DB.Create(&user)

	return c.JSON(user)
}
