package handlers

import (
	"net/http"

	"easy-wallet-be/src/utils"

	"github.com/labstack/echo/v4"
)

func Register(c echo.Context) error {

	return utils.HandleResponse(
		c,
		http.StatusCreated,
		"Register",
		nil,
	)
}
