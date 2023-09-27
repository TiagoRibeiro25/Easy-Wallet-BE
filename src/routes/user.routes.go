package routes

import (
	handlers "easy-wallet-be/src/handlers/user"
	"easy-wallet-be/src/middlewares"

	"github.com/labstack/echo/v4"
)

// UserRoutes sets up the user routes for the server.
func UserRoutes(server *echo.Group) {
	userRoutes := server.Group("/v1/user")

	// Verify user
	userRoutes.PATCH("/verify/:token", handlers.VerifyUser)

	// Get logged in user
	userRoutes.GET("/me", handlers.GetUser, middlewares.ValidateSessionMiddleware)
}
