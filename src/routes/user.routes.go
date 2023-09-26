package routes

import (
	handlers "easy-wallet-be/src/handlers/user"
	"easy-wallet-be/src/middlewares"

	"github.com/labstack/echo/v4"
)

// UserRoutes sets up the user routes for the server.
func UserRoutes(server *echo.Group) {
	userRoutes := server.Group("/v1/user")

	// Use the ValidateSessionMiddleware to check if the user is logged in
	userRoutes.Use(middlewares.ValidateSessionMiddleware)

	userRoutes.GET("", handlers.GetUser)
}
