package handlers

import (
	controllers "easy-wallet-be/src/controllers/auth"
	schemas "easy-wallet-be/src/data/schemas/auth/register"
	"easy-wallet-be/src/models"
	"easy-wallet-be/src/services"
	"easy-wallet-be/src/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Register is a handler function that registers a new user.
// It receives a context and a request body with the user data.
// It checks if there's already a user with the same email, generates tokens, hashes the password,
// calls the controller to register the user, sends a verification email and returns a response with the registered user data.
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

	// Generate tokens
	verifyUserToken, verifyUserTokenErr := utils.GenerateToken()
	resetPasswordToken, resetPasswordTokenErr := utils.GenerateToken()

	// Hash password
	hashedPassword, hashedPasswordErr := utils.HashPassword(bodyData.Password)

	// Check if there was an error while generating tokens or hashing the password
	if verifyUserTokenErr != nil || resetPasswordTokenErr != nil || hashedPasswordErr != nil {
		return handleServerError(c)
	}

	// Call the controller to register the user
	responseData, err := controllers.Register(c, bodyData, verifyUserToken, resetPasswordToken, hashedPassword)
	if err != nil {
		return handleServerError(c)
	}

	// Create the default categories for the user
	if err := models.AddDefaultCategories(responseData.ID); err != nil {
		return handleServerError(c)
	}

	// Send the verification email
	services.SendEmail(
		bodyData.Email,
		bodyData.DisplayName,
		"Easy Wallet - Verify User",
		"<h4>Verify User Token</h4><p>"+verifyUserToken+"</p>",
	)

	return utils.HandleResponse(
		c,
		http.StatusCreated,
		"Successfully registered user",
		responseData,
	)
}

func handleServerError(context echo.Context) error {
	return utils.HandleResponse(
		context,
		http.StatusInternalServerError,
		"Error while registering user",
		nil,
	)
}
