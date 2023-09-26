package handlers

import (
	controllers "easy-wallet-be/src/controllers/cronjob"
	"easy-wallet-be/src/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

// DeleteUnverifiedUsers is a handler function that deletes unverified users from the database.
// It takes an echo.Context object as input and returns an error.
// If the deletion is successful, it returns a success message with HTTP status code 200.
// If there is an error during the deletion process, it returns an error message with HTTP status code 500.
func DeleteUnverifiedUsers(c echo.Context) error {
	responseMessage, err := controllers.DeleteUnverifiedUsers()
	if err != nil {
		return utils.HandleResponse(
			c,
			http.StatusInternalServerError,
			"Error while deleting unverified users",
			nil,
		)
	}

	return utils.HandleResponse(
		c,
		http.StatusOK,
		responseMessage,
		nil,
	)
}
