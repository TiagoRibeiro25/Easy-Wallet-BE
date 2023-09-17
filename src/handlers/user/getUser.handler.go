package handlers

import (
	"net/http"

	"easy-wallet-be/src/utils"

	"github.com/labstack/echo/v4"
)

type User struct {
	ID       int    `json:"id"`
	UserName string `json:"name"`
	Email    string `json:"email"`
}

func GetUser(c echo.Context) error {
	user := &User{
		ID:       1,
		UserName: "John Doe",
		Email:    "john.doe@email.com",
	}

	return utils.HandleResponse(
		c,
		http.StatusOK,
		"User found",
		user,
	)
}
