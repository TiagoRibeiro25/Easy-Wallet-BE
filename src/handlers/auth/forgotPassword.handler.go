package handlers

import (
	controllers "easy-wallet-be/src/controllers/auth"
	schemas "easy-wallet-be/src/data/schemas/auth/forgotPassword"
	"easy-wallet-be/src/data/templates"
	"easy-wallet-be/src/services"
	"easy-wallet-be/src/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

// ForgotPassword generates a reset password token to replace the old one,
// updates the user's reset password token, and sends an email with the token.
func ForgotPassword(c echo.Context) error {
	var bodyData schemas.BodyData
	c.Bind(&bodyData)

	// TODO: Move this code to the end of the resetPassword handler instead

	// Generate the reset password token to replace the old one
	token, err := utils.GenerateToken()
	if err != nil {
		return utils.HandleResponse(
			c,
			http.StatusInternalServerError,
			"Error generating reset password token",
			nil,
		)
	}

	// Update the user's reset password token
	user, err := controllers.UpdateUserPasswordToken(bodyData.Email, token)
	if err != nil {
		return utils.HandleResponse(
			c,
			http.StatusNotFound,
			"User not found",
			nil,
		)
	}

	// Send the email with the token
	services.SendEmail(
		user.Email,
		user.DisplayName,
		"Easy Wallet - Reset password token",
		templates.ForgotPassword(token),
	)

	return utils.HandleResponse(
		c,
		http.StatusOK,
		"Reset password token sent to the specified email address",
		nil,
	)
}
