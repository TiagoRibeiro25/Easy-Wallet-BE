package handlers

import (
	controllers "easy-wallet-be/src/controllers/categories"
	schemas "easy-wallet-be/src/data/schemas/categories/addCategory"
	"easy-wallet-be/src/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

// AddCategory handles the creation of a new category for a given user.
// It expects a JSON payload containing the category name and icon ID.
// The userID is retrieved from the context and converted to uint.
// If the conversion fails or the category creation fails, an error response is returned.
// Otherwise, a success response is returned.
func AddCategory(c echo.Context) error {
	var bodyData schemas.BodyData
	c.Bind(&bodyData)

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

	if err := controllers.AddCategory(userIDuint, bodyData.Name, bodyData.IconID); err != nil {
		return utils.HandleResponse(
			c,
			http.StatusInternalServerError,
			"Something went wrong",
			nil,
		)
	}

	return utils.HandleResponse(
		c,
		http.StatusCreated,
		"Category added successfully",
		nil,
	)
}
