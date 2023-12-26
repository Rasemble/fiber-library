package routes

import (
	"github.com/Rasemble/Api-fiber-library/app/controllers"
	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(a *fiber.App) {
	// Create routes group
	route := a.Group("/api/v1")

	// Routes for GET method
	route.Get("/books", controllers.GetBooks)
	route.Get("/books/:id", controllers.GetBook)
	route.Get("/token/new", controllers.GetNewAccessToken)
}
