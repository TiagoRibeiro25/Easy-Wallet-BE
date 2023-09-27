package routes

import (
	handlers "easy-wallet-be/src/handlers/auth"
	"easy-wallet-be/src/middlewares"

	"github.com/labstack/echo/v4"
)

// AuthRoutes sets up the authentication routes for the server.
func AuthRoutes(server *echo.Group) {
	authRoutes := server.Group("/auth")

	// Register a new user
	authRoutes.POST(
		"/register",
		handlers.Register,
		middlewares.ValidateJSONSchema("auth/register"),
	)

	// Login a user
	authRoutes.POST(
		"/login",
		handlers.Login,
		middlewares.ValidateJSONSchema("auth/login"),
	)
}
