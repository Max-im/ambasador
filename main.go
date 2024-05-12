package main

import (
	"log"
	"shop/src/database"

	"github.com/gofiber/fiber/v3"
)

func main() {
	database.Connect()

	// Initialize a new Fiber app
	app := fiber.New()

	// Define a route for the GET method on the root path '/'
	app.Get("/", func(c fiber.Ctx) error {
		// Send a string response to the client
		return c.SendString("Hello, World777!")
	})

	// Start the server on port 3000
	log.Fatal(app.Listen(":3000"))
}
