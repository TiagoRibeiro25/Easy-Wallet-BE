package handlers

import (
	controllers "easy-wallet-be/src/controllers/categories"
	schemas "easy-wallet-be/src/data/schemas/categories/updateCategory"
	"easy-wallet-be/src/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

// UpdateCategory updates a category with the given ID, name and icon ID.
// It first checks if the category exists and belongs to the user that is trying to edit it.
// If the category exists and belongs to the user, it updates the category and returns a success message.
// If any error occurs during the process, it returns an error response.
func UpdateCategory(c echo.Context) error {
	var bodyData schemas.BodyData
	c.Bind(&bodyData)

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

	// Check if it belongs to the user that is trying to edit it
	if category.UserID != userIDuint {
		return utils.HandleResponse(
			c,
			http.StatusForbidden,
			"You are not allowed to edit this category",
			nil,
		)
	}

	// Update the category
	if err := controllers.UpdateCategory(categoryIDuint, bodyData.Name, bodyData.IconID); err != nil {
		return handleServerError(c)
	}

	return utils.HandleResponse(
		c,
		http.StatusOK,
		"Category updated successfully",
		nil,
	)
}
