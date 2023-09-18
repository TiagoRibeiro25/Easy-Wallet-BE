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

	// Set Middlewares
	color.Cyan("Setting up the middlewares...")

	// Set CORS
	server.Use(middleware.CORSWithConfig(
		configs.GetCorsConfig(),
	))

	// Set Rate Limiter
	server.Use(middleware.RateLimiterWithConfig(
		configs.GetRateLimiterConfig(),
	))

	// Set Logger
	server.Use(middleware.RequestLoggerWithConfig(
		configs.GetLoggerConfig(),
	))

	// Set Security
	server.Use(middleware.SecureWithConfig(
		configs.GetSecurityConfig(),
	))

	// Set Recover
	server.Use(middleware.RecoverWithConfig(
		configs.GetRecoverConfig(),
	))

	// Set Body Limit
	server.Use(middleware.BodyLimitWithConfig(
		configs.GetBodyLimitConfig(),
	))

	// Set Gzip
	server.Use(middleware.GzipWithConfig(
		configs.GetGzipConfig(),
	))

	color.Green("Middlewares are set up")

	// Register all routes
	routes.Init(server)

	return server
}
