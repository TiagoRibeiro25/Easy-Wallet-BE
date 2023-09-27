package handlers

import (
	"easy-wallet-be/src/configs"
	controllers "easy-wallet-be/src/controllers/auth"
	"easy-wallet-be/src/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Logout(c echo.Context) error {
	sessionID, err := utils.ReadCookie(c, configs.GetCookiesConfig().AuthCookieName)
	if err != nil {
		return sendResponse(c)
	}

	// Delete the session from the database
	err = controllers.DeleteSession(sessionID)
	if err != nil {
		// If the session is not found, it means that the user is already logged out
		return sendResponse(c)
	}

	// Delete the session cookie
	utils.DeleteCookie(c, configs.GetCookiesConfig().AuthCookieName)

	return sendResponse(c)
}

func sendResponse(context echo.Context) error {
	return utils.HandleResponse(
		context,
		http.StatusOK,
		"Logged out successfully",
		nil,
	)
}
