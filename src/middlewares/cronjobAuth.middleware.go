package middlewares

import (
	"easy-wallet-be/src/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

// CronjobAuthMiddleware is a middleware function that checks if the Authorization header
// of the incoming request matches the expected authorization key for cronjobs. If the
// authorization key is incorrect, it returns a 401 Unauthorized response. Otherwise, it
// calls the next handler in the chain.
func CronjobAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		expectedAuthKey := utils.GetEnv("CRONJOB_AUTH_KEY")

		if authHeader != expectedAuthKey {
			utils.HandleResponse(
				c,
				http.StatusUnauthorized,
				"Incorrect authorization key",
				nil,
			)
		}

		return next(c)
	}
}
