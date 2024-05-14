package routes

import (
	"shop/src/controllers"

	"github.com/gofiber/fiber/v3"
)

func Setup(app *fiber.App) {
	api := app.Group("api")
	admin := api.Group("admin")

	admin.Post("/register", controllers.Register)
	admin.Post("/login", controllers.Login)
	admin.Get("/me", controllers.User)
	admin.Post("/logout", controllers.Logout)
}
