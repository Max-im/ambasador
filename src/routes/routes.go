package routes

import (
	"shop/src/controllers"
	"shop/src/middlewares"

	"github.com/gofiber/fiber/v3"
)

func Setup(app *fiber.App) {
	api := app.Group("api")
	admin := api.Group("admin")

	admin.Post("/register", controllers.Register)
	admin.Post("/login", controllers.Login)

	adminAuthenticated := admin.Use(middlewares.IsAuthenticated)
	adminAuthenticated.Get("/me", controllers.User)
	adminAuthenticated.Post("/logout", controllers.Logout)
	adminAuthenticated.Put("/user/info", controllers.UpdateInfo)
	adminAuthenticated.Put("/user/password", controllers.UpdatePassword)

	adminAuthenticated.Get("/ambassadors", controllers.Ambassadors)
	adminAuthenticated.Get("/product", controllers.Products)
	adminAuthenticated.Post("/product", controllers.CreateProduct)
	adminAuthenticated.Get("/product/:id", controllers.GetProduct)
	adminAuthenticated.Put("/product/:id", controllers.UpdateProduct)
	adminAuthenticated.Delete("/product/:id", controllers.DeleteProduct)
}
