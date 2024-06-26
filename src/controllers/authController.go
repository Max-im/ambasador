package controllers

import (
	"encoding/json"
	"net/http"
	"shop/src/database"
	"shop/src/middlewares"
	"shop/src/models"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v3"
)

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

type UpdateRequestData struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type UpdatePassRequestData struct {
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirm"`
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
		IsAmbassador: strings.Contains(c.OriginalURL(), "/api/ambassador"),
	}

	user.SetPassword(data.Password)

	database.DB.Create(&user)

	if user.ID == 0 {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message": "User already exists",
		})
	}

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

	isAmbassador := strings.Contains(c.OriginalURL(), "/api/ambassador")
	var scope string

	if isAmbassador {
		scope = "ambassador"
	} else {
		scope = "admin"
	}

	if !isAmbassador && user.IsAmbassador {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	tokenString, err := middlewares.GenerateJWT(user.ID, scope)
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

func User(c fiber.Ctx) error {
	UserID, _ := middlewares.GetUserId(c)

	var user models.User

	database.DB.Where("id = ?", UserID).First(&user)
	if user.ID == 0 {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid credentials",
		})
	}

	if strings.Contains(c.OriginalURL(), "/api/ambassador") {
		ambasador := models.Ambassador(user)
		ambasador.CalculateRevenue(database.DB)
		return c.JSON(ambasador)
	}

	return c.JSON(user)
}

func Logout(c fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "Success",
	})
}

func UpdateInfo(c fiber.Ctx) error {
	var data UpdateRequestData
	err := json.Unmarshal(c.Body(), &data)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body format",
			"error":   err.Error(),
		})
	}

	if data.FirstName == "" || data.LastName == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Missing required fields: 'first_name' or 'last_name'",
		})
	}

	UserID, _ := middlewares.GetUserId(c)

	parsedUserID, err := strconv.ParseUint(UserID, 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid user ID format",
			"error":   err.Error(),
		})
	}

	var currentUser models.User

	database.DB.Where("id = ?", UserID).First(&currentUser)

	user := models.User{
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Email:     currentUser.Email,
	}
	user.ID = uint(parsedUserID)

	database.DB.Model(&user).Updates(&user)

	return c.JSON(user)
}

func UpdatePassword(c fiber.Ctx) error {
	var data UpdatePassRequestData
	err := json.Unmarshal(c.Body(), &data)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body format",
			"error":   err.Error(),
		})
	}

	if data.Password == "" || data.PasswordConfirm == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Missing required fields: 'Password' or 'PasswordConfirm'",
		})
	}

	if data.Password != data.PasswordConfirm {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Passwords dont match",
		})
	}

	UserID, _ := middlewares.GetUserId(c)

	var user models.User

	database.DB.Where("id = ?", UserID).First(&user)

	user.SetPassword(data.Password)

	database.DB.Save(&user)

	return c.JSON(user)
}
