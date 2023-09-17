package routes

import (
	"easy-wallet-be/src/handlers"

	"github.com/fatih/color"
	"github.com/labstack/echo/v4"
)

// Takes the echo server as a parameter to register all routes
func Init(server *echo.Echo) {
	color.Cyan("Registering routes...")

	server.GET("/api", handlers.Home)
	server.GET("*", handlers.NotFound)

	color.Green("Routes registered successfully")
}
