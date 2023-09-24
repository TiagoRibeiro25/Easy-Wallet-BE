package handlers

import (
	controllers "easy-wallet-be/src/controllers/auth"
	schemas "easy-wallet-be/src/data/schemas/auth/register"
	"easy-wallet-be/src/models"
	"easy-wallet-be/src/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Register(c echo.Context) error {
	var bodyData schemas.BodyData
	c.Bind(&bodyData)

	db := models.DB()

	// Check if there's already an user with the same email
	user := db.Where("email = ?", bodyData.Email).First(&models.User{})
	if user.RowsAffected > 0 {
		return utils.HandleResponse(
			c,
			http.StatusConflict,
			"The email is already in use",
			nil,
		)
	}

	// Create the account
	controllers.Register(c, bodyData)

	return nil
}