package configs

import (
	"github.com/labstack/echo/v4/middleware"
)

func GetBodyLimitConfig() middleware.BodyLimitConfig {
	return middleware.BodyLimitConfig{
		Skipper: middleware.DefaultSkipper,
		Limit:   "2M",
	}
}
