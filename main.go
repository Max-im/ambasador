package main

import (
	"log"
	"shop/src/database"
	"shop/src/routes"

	"github.com/gofiber/fiber/v3"
)

func main() {
	if err := database.Connect(); err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}

	// Initialize a new Fiber app
	app := fiber.New()

	// Routes
	routes.Setup(app)

	// Start the server on port 3000
	log.Fatal(app.Listen(":3000"))
}
