package handlers

import (
	"net/http"

	"easy-wallet-be/src/utils"

	"github.com/labstack/echo/v4"
)

func Home(c echo.Context) error {
	return utils.HandleResponse(
		c,
		http.StatusOK,
		"Welcome to the Easy Wallet API",
		nil,
	)
}

func NotFound(c echo.Context) error {
	return utils.HandleResponse(
		c,
		http.StatusNotFound,
		"Route not found",
		nil,
	)
}
