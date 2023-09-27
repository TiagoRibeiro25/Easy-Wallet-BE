package handlers

import (
	"net/http"

	controllers "easy-wallet-be/src/controllers/user"
	"easy-wallet-be/src/utils"

	"github.com/labstack/echo/v4"
)

func handleServerError(c echo.Context, err error) error {
	return utils.HandleResponse(
		c,
		http.StatusInternalServerError,
		"Something went wrong on our end",
		nil,
	)
}

// GetUser returns a user with dummy data.
func GetUser(c echo.Context) error {
	// Get the userID from the context (added by the ValidateSessionMiddleware)
	userID := utils.ConvertVarToString(c.Get("userID"))

	// Convert the userID to uint
	userIDuint, err := utils.ConvertStringToUint(userID)
	if err != nil {
		return handleServerError(c, err)
	}

	// Get the user from the database
	user, err := controllers.GetUser(userIDuint)
	if err != nil {
		return handleServerError(c, err)
	}

	return utils.HandleResponse(
		c,
		http.StatusOK,
		"User found",
		user,
	)
}
