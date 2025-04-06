package routes

import "github.com/gofiber/fiber/v2"

// welcome is a simple handler function for the welcome endpoint.
func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to an Awesome API")
}

// ApiRoutes sets up the routes for the API with versioning.
func ApiRoutes(app *fiber.App) {
	// Versioned API Route
	api := app.Group("/api")

	// Welcome endpoint
	api.Get("/", welcome)
}
