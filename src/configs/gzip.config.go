package configs

import "github.com/labstack/echo/v4/middleware"

func GetGzipConfig() middleware.GzipConfig {
	return middleware.GzipConfig{
		Skipper: middleware.DefaultSkipper,
		Level:   5,
	}
}
