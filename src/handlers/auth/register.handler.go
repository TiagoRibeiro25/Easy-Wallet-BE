package handlers

import (
	controllers "easy-wallet-be/src/controllers/auth"
	schemas "easy-wallet-be/src/data/schemas/auth/register"
	"easy-wallet-be/src/services"
	"easy-wallet-be/src/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Register is a handler function that receives a request context and registers a new user account.
// It checks if there's already a user with the same email and returns an error if so.
// Otherwise, it creates the account and returns a response.
func Register(c echo.Context) error {
	var bodyData schemas.BodyData
	c.Bind(&bodyData)

	// Check if there's already a user with the same email
	if controllers.DoesUserExist(bodyData.Email) {
		return utils.HandleResponse(
			c,
			http.StatusConflict,
			"The email is already in use",
			nil,
		)
	}

	// Call the controller to register the user
	responseData, verifyUserToken, err := controllers.Register(c, bodyData)
	if err != nil {
		return utils.HandleResponse(
			c,
			http.StatusInternalServerError,
			"Error while registering user",
			nil,
		)
	}

	// Send the verification email
	services.SendEmail(
		bodyData.Email,
		bodyData.DisplayName,
		"Welcome to Easy Wallet",
		"<h4>Verify User Token</h4><p>"+verifyUserToken+"</p>",
	)

	return utils.HandleResponse(
		c,
		http.StatusCreated,
		"Successfully registered user",
		responseData,
	)
}
