package routes

import (
	handlers "easy-wallet-be/src/handlers/cronjob"
	"easy-wallet-be/src/middlewares"

	"github.com/labstack/echo/v4"
)

//TODO: Add a cron job to delete the token after 15 minutes and create a new one
//TODO: Add a cron job to delete the user after 24 hours if the user is not verified

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
