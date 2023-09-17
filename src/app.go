package main

import (
	"easy-wallet-be/src/configs"
	"easy-wallet-be/src/routes"

	"github.com/fatih/color"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// The function creates an HTTP server using the Echo framework
func Server() *echo.Echo {
	color.Blue("Setting up the server...")

	server := echo.New()

	// Set CORS
	server.Use(middleware.CORSWithConfig(
		configs.GetCorsConfig(),
	))

	// Set Rate Limiter
	server.Use(middleware.RateLimiterWithConfig(
		configs.GetRateLimiterConfig(),
	))

	// Register all routes
	routes.Init(server)

	return server
}
