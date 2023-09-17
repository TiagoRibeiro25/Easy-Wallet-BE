package configs

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"easy-wallet-be/src/utils"
)

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
