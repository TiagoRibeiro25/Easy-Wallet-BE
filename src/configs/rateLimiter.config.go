package configs

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"easy-wallet-be/src/utils"
)

// GetRateLimiterConfig returns a middleware.RateLimiterConfig struct that can be used to limit the rate of incoming requests.
// The rate limiter is configured to allow 3 requests per minute, with a burst of 3 requests.
// The identifier extractor function extracts the client's IP address as the identifier.
// The error handler function returns a 403 Forbidden response when the rate limit is exceeded.
// The deny handler function returns a 429 Too Many Requests response when the rate limit is exceeded.
func GetRateLimiterConfig() middleware.RateLimiterConfig {
	return middleware.RateLimiterConfig{
		Skipper: middleware.DefaultSkipper,
		Store: middleware.NewRateLimiterMemoryStoreWithConfig(
			middleware.RateLimiterMemoryStoreConfig{Rate: 3, Burst: 3, ExpiresIn: time.Minute},
		),
		IdentifierExtractor: func(c echo.Context) (string, error) {
			id := c.RealIP()
			return id, nil
		},
		ErrorHandler: func(c echo.Context, err error) error {
			return utils.HandleResponse(
				c,
				http.StatusForbidden,
				"Too many requests",
				nil,
			)
		},
		DenyHandler: func(c echo.Context, identifier string, err error) error {
			return utils.HandleResponse(
				c,
				http.StatusTooManyRequests,
				"Too many requests",
				nil,
			)
		},
	}
}
