package handlers

import (
	controllers "easy-wallet-be/src/controllers/cronjob"
	"easy-wallet-be/src/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func DeleteUnverifiedUsers(c echo.Context) error {
	responseMessage, err := controllers.DeleteUnverifiedUsers(c)
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
