package main

import (
    "log"
    "github.com/gofiber/fiber/v3"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

func main() {
    dsn := "host=db user=postgres password=postgres dbname=shop port=5432 sslmode=disable TimeZone=Asia/Shanghai"
    _, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    
    if err != nil {
        log.Fatal("failed to connect database:", err)
    }

    // Initialize a new Fiber app
    app := fiber.New()

    // Define a route for the GET method on the root path '/'
    app.Get("/", func(c fiber.Ctx) error {
        // Send a string response to the client
        return c.SendString("Hello, World 👋!")
    })

    // Start the server on port 3000
    log.Fatal(app.Listen(":3000"))
}