package main

import (
	configs "github.com/Rasemble/Api-fiber-library/pkg/config"
	"github.com/Rasemble/Api-fiber-library/pkg/middleware"
	"github.com/Rasemble/Api-fiber-library/pkg/routes"
	"github.com/Rasemble/Api-fiber-library/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Define Fiber config
	config := configs.FiberConfig()

	// Define new fiber app with config
	app := fiber.New(config)

	// Register meddleware
	middleware.FiberMiddleware(app)

	// Register routes
	routes.PrivatesRoutes(app)
	routes.PublicRoutes(app)
	routes.NotFoundRoute(app)

	utils.StartServer(app)

}
