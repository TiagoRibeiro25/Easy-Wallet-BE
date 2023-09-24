package routes

import (
	handlers "easy-wallet-be/src/handlers/user"

	"github.com/labstack/echo/v4"
)

// UserRoutes sets up the user routes for the server.
func UserRoutes(server *echo.Group) {
	userRoutes := server.Group("/v1/user")

	userRoutes.GET("", handlers.GetUser)
}
