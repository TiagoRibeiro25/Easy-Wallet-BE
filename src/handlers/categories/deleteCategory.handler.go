package handlers

import (
	controllers "easy-wallet-be/src/controllers/categories"
	"easy-wallet-be/src/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func handleServerError(context echo.Context) error {
	return utils.HandleResponse(
		context,
		http.StatusInternalServerError,
		"Something went wrong",
		nil,
	)
}

// DeleteCategory handles the DELETE request to delete a category by ID.
// It receives the category ID from the request URL parameter and the user ID from the context.
// It checks if the category exists and belongs to the user that is trying to delete it.
// If the category exists and belongs to the user, it deletes the category and returns a success message.
// If any error occurs, it returns an error response.
func DeleteCategory(c echo.Context) error {
	// Get the categoryID from the request (url param)
	categoryID := c.Param("id")
	// Get the userID from the context (added by the ValidateSessionMiddleware)
	userID := utils.ConvertVarToString(c.Get("userID"))

	// Convert the categoryID to uint
	categoryIDuint, err := utils.ConvertStringToUint(categoryID)
	if err != nil {
		return handleServerError(c)
	}

	// Check if the category exists
	category, err := controllers.GetCategory(categoryIDuint)
	if err != nil {
		return utils.HandleResponse(
			c,
			http.StatusNotFound,
			"Category not found",
			nil,
		)
	}

	// Convert the userID to uint
	userIDuint, err := utils.ConvertStringToUint(userID)
	if err != nil {
		return handleServerError(c)
	}

	// Check if it belongs to the user that is trying to delete it
	if category.UserID != userIDuint {
		return utils.HandleResponse(
			c,
			http.StatusForbidden,
			"You are not allowed to delete this category",
			nil,
		)
	}

	// Delete the category
	err = controllers.DeleteCategory(categoryIDuint)
	if err != nil {
		return handleServerError(c)
	}

	return utils.HandleResponse(
		c,
		http.StatusOK,
		"Category deleted successfully",
		nil,
	)
}
