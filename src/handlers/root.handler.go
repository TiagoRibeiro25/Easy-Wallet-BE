package handlers

import (
	"net/http"

	"easy-wallet-be/src/utils"

	"github.com/labstack/echo/v4"
)

// Home is a handler function that returns a welcome message
func Home(c echo.Context) error {
	return utils.HandleResponse(
		c,
		http.StatusOK,
		"Welcome to the Easy Wallet API",
		nil,
	)
}

// NotFound is a handler function that returns a "Route not found" message
func NotFound(c echo.Context) error {
	return utils.HandleResponse(
		c,
		http.StatusNotFound,
		"Route not found",
		nil,
	)
}
