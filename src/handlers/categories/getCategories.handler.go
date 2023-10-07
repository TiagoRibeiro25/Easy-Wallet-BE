package handlers

import (
	controllers "easy-wallet-be/src/controllers/categories"
	"easy-wallet-be/src/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetCategories retrieves all categories for a given user ID.
// It expects a valid userID in the context, added by the ValidateSessionMiddleware.
// It returns a JSON response with the categories data and a success message.
// If the userID is invalid or the user is not found, it returns an error message with a 404 status code.
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
