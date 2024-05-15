package main

import (
	"log"
	"shop/src/database"
	"shop/src/routes"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func main() {
	if err := database.Connect(); err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}

	database.SetupRedis()

	// Initialize a new Fiber app
	app := fiber.New()

	// Cors
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowCredentials: true,
	}))
	// app.Use(cors.New())

	// Routes
	routes.Setup(app)

	// Start the server on port 3000
	log.Fatal(app.Listen(":3000"))
}
