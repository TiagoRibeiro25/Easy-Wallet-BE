package handlers

import (
	controllers "easy-wallet-be/src/controllers/auth"
	schemas "easy-wallet-be/src/data/schemas/auth/login"
	"easy-wallet-be/src/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

const SELECT_QUERY = "users.id, users.email, users.display_name, users.user_verified, users.verify_user_token ,users.currency, users.created_at, passwords.password"

// Login handles the login request for the user. It receives the user's email, password and rememberMe, fetches the user by email,
// compares the password, checks if the user is verified, creates a session, and returns a response with the user's
// information and a session cookie. If the user is not found, the password is incorrect, or the user is not verified,
// it returns an error response.
func Login(c echo.Context) error {
	var bodyData schemas.BodyData
	if err := c.Bind(&bodyData); err != nil {
		utils.HandleResponse(
			c,
			http.StatusBadRequest,
			"Invalid request data",
			nil,
		)
	}

	// Find the user by email
	user, err := controllers.FetchUserByEmail(c, bodyData.Email)
	if err != nil {
		return utils.HandleResponse(
			c,
			http.StatusNotFound,
			"User not found",
			nil,
		)
	}

	// Compare the password
	if !(utils.ComparePassword(user.Password, bodyData.Password)) {
		return utils.HandleResponse(
			c,
			http.StatusUnauthorized,
			"Invalid credentials",
			nil,
		)
	}

	// Convert the user's ID to uint
	userIDuint, err := utils.ConvertStringToUint(user.ID)
	if err != nil {
		return utils.HandleResponse(
			c,
			http.StatusInternalServerError,
			"An error occurred while trying to log in",
			nil,
		)
	}

	// Create a session
	sessionID, err := controllers.CreateSession(userIDuint, bodyData.RememberMe)
	if err != nil {
		return utils.HandleResponse(
			c,
			http.StatusInternalServerError,
			"An error occurred while trying to create a session",
			nil,
		)
	}

	// Create a session cookie
	utils.WriteCookie(
		c,
		"easywallet-session-id",
		sessionID,
		utils.GetCookieExpiration(bodyData.RememberMe),
	)

	return utils.HandleResponse(
		c,
		http.StatusCreated,
		"Successfully logged in",
		schemas.ResponseData{
			ID:           userIDuint,
			Email:        user.Email,
			DisplayName:  user.DisplayName,
			UserVerified: user.UserVerified,
			Currency:     user.Currency,
			CreatedAt:    user.CreatedAt,
		},
	)
}