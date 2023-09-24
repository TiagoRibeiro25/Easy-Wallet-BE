package configs

import "github.com/labstack/echo/v4/middleware"

// GetGzipConfig returns the GzipConfig struct with default values.
// The Skipper field is set to middleware.DefaultSkipper and the Level field is set to 5.
func GetGzipConfig() middleware.GzipConfig {
	return middleware.GzipConfig{
		Skipper: middleware.DefaultSkipper,
		Level:   5,
	}
}
