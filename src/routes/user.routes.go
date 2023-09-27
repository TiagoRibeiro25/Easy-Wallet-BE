package routes

import (
	handlers "easy-wallet-be/src/handlers/user"
	"easy-wallet-be/src/middlewares"

	"github.com/labstack/echo/v4"
)

// UserRoutes sets up the user routes for the server.
func UserRoutes(server *echo.Group) {
	userRoutes := server.Group("/v1/user")

	// Get logged in user (not finished yet)
	userRoutes.GET("", handlers.GetUser, middlewares.ValidateSessionMiddleware)

	// Verify user
	userRoutes.PATCH("/verify/:token", handlers.VerifyUser)
}
