package handlers

import (
	categoriesControllers "easy-wallet-be/src/controllers/categories"
	expensesControllers "easy-wallet-be/src/controllers/expenses"
	schemas "easy-wallet-be/src/data/schemas/expenses/addExpense"
	"easy-wallet-be/src/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

// AddExpense handles the request to add a new expense.
// It takes the request body data from the context and validates it.
// It also checks if the category exists and belongs to the user.
// If all validations pass, it adds the expense and returns the newly created expense.
// If any error occurs during the process, it returns an appropriate error response.
func AddExpense(c echo.Context) error {
	var bodyData schemas.BodyData
	c.Bind(&bodyData)

	// Take the user ID added to the context by the ValidateSessionMiddleware and convert it to a uint
	userIDuint, err := utils.ConvertStringToUint(utils.ConvertVarToString(c.Get("userID")))
	if err != nil {
		return utils.HandleResponse(c, http.StatusInternalServerError, "Something went wrong", nil)
	}

	// Validate the date
	if !utils.IsDateValid(bodyData.Date) || utils.IsDateInFuture(bodyData.Date) {
		return utils.HandleResponse(c, http.StatusBadRequest, "Invalid date", nil)
	}

	// Check if the category exists and it's from the user
	category, err := categoriesControllers.GetCategory(bodyData.CategoryID)
	if err != nil || category.UserID != userIDuint {
		return utils.HandleResponse(c, http.StatusBadRequest, "Invalid category", nil)
	}

	newExpense, err := expensesControllers.AddExpense(userIDuint, expensesControllers.NewExpense{
		Name:        bodyData.Name,
		Cost:        bodyData.Cost,
		Date:        bodyData.Date,
		Description: bodyData.Description,
		CategoryID:  bodyData.CategoryID,
	})

	if err != nil {
		return utils.HandleResponse(c, http.StatusInternalServerError, "Something went wrong", nil)
	}

	return utils.HandleResponse(c, http.StatusCreated, "Expense added successfully", newExpense)
}
