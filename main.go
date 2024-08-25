package main

import (
	"bfhl/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Initialize Fiber
	app := fiber.New()

	// Define routes
	app.Get("/bfhl", routes.GetUser)
	app.Post("/bfhl", routes.PostUser)

	// Start the server
	app.Listen(":3000")
}
