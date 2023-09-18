package main

import (
	"easy-wallet-be/src/middlewares"
	"easy-wallet-be/src/routes"

	"github.com/fatih/color"
	"github.com/labstack/echo/v4"
)

// The function creates an HTTP server using the Echo framework
func Server() *echo.Echo {
	color.Blue("Setting up the server...")

	server := echo.New()

	// Set Middlewares
	middlewares.SetUpMiddlewares(server)

	// Register all routes
	routes.Init(server)

	return server
}
