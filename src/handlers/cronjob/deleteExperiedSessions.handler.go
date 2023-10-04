package handlers

import (
	controllers "easy-wallet-be/src/controllers/cronjob"
	"easy-wallet-be/src/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// DeleteExpiredSessions is a handler function that deletes expired sessions.
// It calls the DeleteExpiredSessions function from the controllers package to perform the deletion.
// If an error occurs during the deletion process, it returns an HTTP 500 response with an error message.
// Otherwise, it returns an HTTP 200 response with a success message.
func DeleteExpiredSessions(c echo.Context) error {
	sessionsDeleted, err := controllers.DeleteExpiredSessions()
	if err != nil {
		return utils.HandleResponse(
			c,
			http.StatusInternalServerError,
			"Error while deleting expired sessions",
			nil,
		)
	}

	return utils.HandleResponse(
		c,
		http.StatusOK,
		"Successfully deleted "+strconv.Itoa(sessionsDeleted)+" expired sessions",
		nil,
	)
}
