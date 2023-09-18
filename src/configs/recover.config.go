package configs

import (
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func GetRecoverConfig() middleware.RecoverConfig {
	return middleware.RecoverConfig{
		StackSize:         1 << 10, // 1 KB
		DisableStackAll:   false,
		DisablePrintStack: false,
		LogLevel:          log.ERROR,
	}
}
