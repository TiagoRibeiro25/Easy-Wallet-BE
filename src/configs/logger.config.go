package configs

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func getStatusColor(statusCode int) color.Attribute {
	switch {
	case statusCode >= 0 && statusCode < 100:
		return color.FgHiMagenta
	case statusCode >= 100 && statusCode < 200:
		return color.FgHiBlue
	case statusCode >= 200 && statusCode < 300:
		return color.FgHiGreen
	case statusCode >= 300 && statusCode < 400:
		return color.FgHiCyan
	case statusCode >= 400 && statusCode < 500:
		return color.FgHiYellow
	case statusCode >= 500 && statusCode < 600:
		return color.FgHiRed
	default:
		return color.FgHiWhite
	}
}

func getMethodColor(method string) color.Attribute {
	switch method {
	case "GET":
		return color.FgHiGreen
	case "POST":
		return color.FgHiYellow
	case "PUT":
		return color.FgHiBlue
	case "DELETE":
		return color.FgHiRed
	case "PATCH":
		return color.FgHiMagenta
	case "HEAD":
		return color.FgHiCyan
	case "OPTIONS":
		return color.FgHiWhite
	default:
		return color.FgHiWhite
	}
}

func GetLoggerConfig() middleware.RequestLoggerConfig {
	return middleware.RequestLoggerConfig{
		LogStatus:  true,
		LogURI:     true,
		LogMethod:  true,
		LogLatency: true,
		BeforeNextFunc: func(c echo.Context) {
			c.Set("customValueFromContext", 42)
		},
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {

			// Date: [dd-mm-yyyy hh:mm:ss:ms]
			color.New(color.FgHiBlue).Printf("[%s] - ",
				v.StartTime.Format("02-01-2006 15:04:05.000"),
			)

			color.New(color.FgHiBlue).Print("REQUEST: ")

			// Method
			color.New(color.FgHiCyan).Print("method: ")
			color.New(getMethodColor(v.Method)).Print(v.Method)

			// Status
			color.New(color.FgHiCyan).Print(", status: ")
			color.New(getStatusColor(v.Status)).Print(v.Status)

			// URI
			color.New(color.FgHiCyan).Print(", uri: ")
			color.New(color.FgHiYellow).Print(v.URI)

			fmt.Println()

			return nil
		},
	}
}
