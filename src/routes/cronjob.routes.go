package routes

import (
	handlers "easy-wallet-be/src/handlers/cronjob"
	"easy-wallet-be/src/middlewares"

	"github.com/labstack/echo/v4"
)

// CronjobRoutes sets up routes that are used by cronjobs.
func CronjobRoutes(server *echo.Group) {
	cronjobRoutes := server.Group("/cronjob")

	// Use the CronjobAuthMiddleware to check if the Authorization header of the incoming
	cronjobRoutes.Use(middlewares.CronjobAuthMiddleware)

	// Delete the user account if the user didn't verify their email address within 24 hours
	cronjobRoutes.DELETE(
		"/delete-unverified-users",
		handlers.DeleteUnverifiedUsers,
	)
}
