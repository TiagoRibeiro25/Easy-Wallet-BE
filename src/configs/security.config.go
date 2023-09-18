package configs

import "github.com/labstack/echo/v4/middleware"

func GetSecurityConfig() middleware.SecureConfig {
	return middleware.SecureConfig{
		XSSProtection:         "1; mode=block",
		ContentTypeNosniff:    "nosniff",
		XFrameOptions:         "SAMEORIGIN",
		HSTSMaxAge:            3600,
		ContentSecurityPolicy: "default-src 'self'",
	}
}
