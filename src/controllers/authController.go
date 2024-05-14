package controllers

import (
	"encoding/json"
	"net/http"
	"shop/src/database"
	"shop/src/models"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

var SECRET_OR_KEY = "your_secret_key"

type RegisterRequestData struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirm"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
}

type LoginRequestData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CustomClaims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
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

	user := models.User{
		Email:        data.Email,
		FirstName:    data.FirstName,
		LastName:     data.LastName,
		IsAmbassador: false,
	}

	user.SetPassword(data.Password)

	database.DB.Create(&user)

	return c.JSON(user)
}

func Login(c fiber.Ctx) error {

	var data LoginRequestData
	err := json.Unmarshal(c.Body(), &data)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body format",
			"error":   err.Error(),
		})
	}

	if data.Email == "" || data.Password == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Missing required fields: 'email' or 'password'",
		})
	}

	var user models.User

	database.DB.Where("email = ?", data.Email).First(&user)

	if user.ID == 0 {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid credentials",
		})
	}

	if err := user.ComparePassword(data.Password); err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid credentials",
		})
	}

	claims := &CustomClaims{
		string(user.ID),
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(SECRET_OR_KEY))
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid credentials",
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    tokenString,
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "Success",
		"user":    user,
	})
}
