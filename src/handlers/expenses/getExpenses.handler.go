package handlers

import (
	controllers "easy-wallet-be/src/controllers/expenses"
	schemas "easy-wallet-be/src/data/schemas/expenses/getExpenses"
	"easy-wallet-be/src/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetExpenses(c echo.Context) error {
	date := c.QueryParam("date")
	page := c.QueryParam("page")
	limit := c.QueryParam("limit")

	if !utils.IsDateValid(date) || utils.IsDateInFuture(date) {
		return utils.HandleResponse(c, http.StatusBadRequest, "Invalid date", nil)
	}

	pageInt, limitInt, arePageAndLimitValid := utils.ArePageAndLimitFromQueryValid(page, limit)
	if !arePageAndLimitValid {
		return utils.HandleResponse(c, http.StatusBadRequest, "Invalid page or limit", nil)
	}

	// Take the user ID added to the context by the ValidateSessionMiddleware and convert it to a uint
	userIDuint, err := utils.ConvertStringToUint(utils.ConvertVarToString(c.Get("userID")))
	if err != nil {
		return utils.HandleResponse(c, http.StatusInternalServerError, "Something went wrong", nil)
	}

	expenses, err := controllers.GetExpenses(userIDuint, date, pageInt, limitInt)
	if err != nil {
		return utils.HandleResponse(c, http.StatusInternalServerError, "Something went wrong", nil)
	}

	// Convert expenses to response data
	var responseExpenses []schemas.ResponseData
	for _, expense := range expenses {
		responseExpenses = append(responseExpenses, schemas.ResponseData{
			ID:          expense.ID,
			Name:        expense.Name,
			Cost:        expense.Cost,
			Date:        expense.Date,
			Description: expense.Description,
			CategoryID:  expense.CategoryID,
			CreatedAt:   expense.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:   expense.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return utils.HandleResponse(c, http.StatusOK, "Expenses retrieved successfully", responseExpenses)
}
