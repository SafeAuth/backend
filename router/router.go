package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// ToDo: add routes

// ConnectRouter is a function that connects all routes to the router
func ConnectRouter(app *fiber.App) {
	// Middleware
	app.Use(logger.New())

	// Routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
}
