package configs

import "github.com/labstack/echo/v4/middleware"

// GetBodyLimitConfig returns the configuration for the body size limit middleware.
// It returns a middleware.BodyLimitConfig struct with the default skipper and a limit of 2 megabytes.
func GetBodyLimitConfig() middleware.BodyLimitConfig {
	return middleware.BodyLimitConfig{
		Skipper: middleware.DefaultSkipper,
		Limit:   "2M",
	}
}
