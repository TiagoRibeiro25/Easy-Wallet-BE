package handlers

import (
	controllers "easy-wallet-be/src/controllers/auth"
	schemas "easy-wallet-be/src/data/schemas/auth/forgotPassword"
	"easy-wallet-be/src/services"
	"easy-wallet-be/src/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

// ForgotPassword handles the forgot password feature by getting the reset password token for the user's email,
// sending an email with the token, and returning a success response if the email was sent successfully.
// It takes in an echo context and returns an error.
func ForgotPassword(c echo.Context) error {
	var bodyData schemas.BodyData
	c.Bind(&bodyData)

	// Get the reset password token
	token, err := controllers.GetResetPasswordTokenByEmail(bodyData.Email)
	if err != nil {
		return utils.HandleResponse(
			c,
			http.StatusNotFound,
			"User with the specified email address not found",
			nil,
		)
	}

	// Send the email with the token
	services.SendEmail(
		bodyData.Email,
		bodyData.Email, // TODO: Change this to the user's display name
		"Easy Wallet - Reset password token",
		"<h4>Update password token</h4><p>"+token+"</p>",
	)

	return utils.HandleResponse(
		c,
		http.StatusOK,
		"Reset password token sent to the specified email address",
		nil,
	)
}
