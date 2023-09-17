package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// define the json response for the root route
type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func Home(c echo.Context) error {
	return c.JSON(http.StatusOK, Response{
		Success: true,
		Message: "Welcome to the Easy Wallet API",
	})
}

func NotFound(c echo.Context) error {
	return c.JSON(http.StatusNotFound, Response{
		Success: false,
		Message: "Route not found",
	})
}
