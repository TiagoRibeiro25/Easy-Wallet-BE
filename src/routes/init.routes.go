package routes

import (
	"easy-wallet-be/src/handlers"

	"github.com/fatih/color"
	"github.com/labstack/echo/v4"
)

// Takes the echo server as a parameter to register all routes
func Init(server *echo.Echo) {
	color.Cyan("Registering routes...")

	// Not found route
	server.GET("*", handlers.NotFound)

	api := server.Group("/api")

	// Home route
	api.GET("", handlers.Home)

	// Register all routes
	UserRoutes(api)
	AuthRoutes(api)
	CategoriesRoutes(api)
	CronjobRoutes(api)

	color.Green("Routes registered successfully")
}
