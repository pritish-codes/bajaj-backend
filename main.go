package main

import (
	"bfhl/routes"

	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// Initialize Fiber
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // Allow all origins, change as per your need
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	// Define routes
	app.Get("/bfhl", routes.GetUser)
	app.Post("/bfhl", routes.PostUser)

	// Start the server
	app.Listen(":3000")
}
