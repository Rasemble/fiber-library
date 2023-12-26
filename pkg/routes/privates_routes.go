package routes

import (
	"github.com/Rasemble/Api-fiber-library/app/controllers"
	"github.com/Rasemble/Api-fiber-library/pkg/middleware"
	"github.com/gofiber/fiber/v2"
)

func PrivatesRoutes(a *fiber.App) {
	route := a.Group("/api/v1")

	// Route for POST method
	route.Post("/book", middleware.JWTProtected(), controllers.CreateBook)

	// Route for Put method
	route.Put("/book", middleware.JWTProtected(), controllers.UpdateBook)

	// Route for Delete method
	route.Delete("/book", middleware.JWTProtected(), controllers.DeleteBook)
}
