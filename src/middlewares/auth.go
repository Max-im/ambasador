package middlewares

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

const SecretKey = "your_secret_key"

type CustomClaims struct {
	UserID string `json:"id"`
	Scope  string `json:"scope"`
	jwt.RegisteredClaims
}

func IsCorrectScope(url string, scope string) bool {
	isAmbassador := strings.Contains(url, "/api/ambassador")

	return (scope == "admin" && !isAmbassador) || (scope == "ambassador" && isAmbassador)
}

func IsAuthenticated(c fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil || !token.Valid {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	payload := token.Claims.(*CustomClaims)
	if !IsCorrectScope(c.OriginalURL(), payload.Scope) {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	return c.Next()
}

func GetUserId(c fiber.Ctx) (string, error) {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil || !token.Valid {
		return "", err
	}

	claims := token.Claims.(*CustomClaims)

	return claims.UserID, nil
}

func GenerateJWT(id uint, scope string) (string, error) {
	claims := &CustomClaims{
		UserID: fmt.Sprintf("%d", id),
		Scope:  scope,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(SecretKey))
}
