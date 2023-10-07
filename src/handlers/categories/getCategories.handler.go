package handlers

import (
	controllers "easy-wallet-be/src/controllers/categories"
	"easy-wallet-be/src/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetCategories returns the categories of the user.
func GetCategories(c echo.Context) error {
	// Get the userID from the context (added by the ValidateSessionMiddleware)
	userID := utils.ConvertVarToString(c.Get("userID"))

	// Convert the userID to uint
	userIDuint, err := utils.ConvertStringToUint(userID)
	if err != nil {
		return utils.HandleResponse(
			c,
			http.StatusInternalServerError,
			"Something went wrong",
			nil,
		)
	}

	responseData, err := controllers.GetCategories(userIDuint)
	if err != nil {
		return utils.HandleResponse(
			c,
			http.StatusNotFound,
			"User not found",
			nil,
		)
	}

	return utils.HandleResponse(
		c,
		http.StatusOK,
		"Categories retrieved successfully",
		responseData,
	)
}
