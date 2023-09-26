package middlewares

import (
	"easy-wallet-be/src/configs"
	controllers "easy-wallet-be/src/controllers/auth"
	"easy-wallet-be/src/utils"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

// TODO: Maybe use a different database to store sessions instead of the main database like Redis or something.

// ValidateSessionMiddleware is a middleware function that checks if the user's session is valid and not expired.
// It takes an echo.HandlerFunc as input and returns an echo.HandlerFunc.
// If the session is invalid or expired, it returns an error response with status code 401 (Unauthorized).
// If the session is valid, it calls the next middleware function.
// If the session expires in less than 30 minutes, it deletes the session and creates a new one.
// It also creates a new session cookie if necessary.
func ValidateSessionMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sessionID, err := utils.ReadCookie(c, configs.GetCookiesConfig().AuthCookieName)
		if err != nil {
			return handleUnauthorized(c)
		}

		// Get the session from the database
		session, err := controllers.GetSession(sessionID)
		if err != nil {
			utils.DeleteCookie(c, configs.GetCookiesConfig().AuthCookieName)
			return handleUnauthorized(c)
		}

		// Check if the session is expired
		if session.ExpiresAt.Before(time.Now()) {
			controllers.DeleteSession(sessionID) // Delete the session from the database
			return handleUnauthorized(c)
		}

		// If the session expires in less than 30 minutes, delete the session and create a new one
		if session.ExpiresAt.Before(time.Now().Add(30 * time.Minute)) {
			controllers.DeleteSession(sessionID)
			newSessionID, err := controllers.CreateSession(session.UserID, session.RememberMe)
			if err != nil {
				return handleUnauthorized(c)
			}

			// Create a session cookie
			utils.WriteCookie(
				c,
				configs.GetCookiesConfig().AuthCookieName,
				newSessionID,
				utils.GetCookieExpiration(session.RememberMe),
			)
		}

		return next(c)
	}
}

func handleUnauthorized(context echo.Context) error {
	return utils.HandleResponse(
		context,
		http.StatusUnauthorized,
		"Unauthorized",
		nil,
	)
}
