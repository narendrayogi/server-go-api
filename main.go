package main

import (
    "log"

    "github.com/gofiber/fiber/v3"
    "github.com/narendrayogi/server-go-api/handler"
)

func main() {
    // Initialize a new Fiber app
    app := fiber.New()

    // Define a route for the GET method on the root path '/'
    app.Get("/", func(c fiber.Ctx) error {
        // Send a string response to the client
        return c.SendString("Hello, World 👋!")
    })

		api := app.Group("/api")      // /api

		v1 := api.Group("/v1")        // /api/v1
		v1.Get("/product", handler.ProductHandler)      // /api/v1/product
		v1.Get("/user", handler.UserHandler)      // /api/v1/user

    // Start the server on port 3000
    log.Fatal(app.Listen(":3000"))
}