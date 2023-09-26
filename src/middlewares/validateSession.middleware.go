package middlewares

import (
	controllers "easy-wallet-be/src/controllers/auth"
	"easy-wallet-be/src/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

// ValidateSessionMiddleware is a middleware function that checks if the session ID stored in the cookie is valid.
// It takes in the next echo.HandlerFunc as a parameter and returns an echo.HandlerFunc.
// If the session ID is not found or is invalid, it returns an unauthorized response.
// If the session ID is valid, it calls the next handler function.
func ValidateSessionMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sessionID, err := utils.ReadCookie(c, "easywallet-session-id")

		if err != nil || !controllers.ValidateSession(sessionID) {
			return utils.HandleResponse(
				c,
				http.StatusUnauthorized,
				"Unauthorized",
				nil,
			)
		}

		return next(c)
	}
}
