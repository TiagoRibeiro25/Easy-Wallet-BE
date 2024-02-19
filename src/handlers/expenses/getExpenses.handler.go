package handlers

import (
	controllers "easy-wallet-be/src/controllers/expenses"
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

	return utils.HandleResponse(c, http.StatusOK, "Expenses retrieved successfully", expenses)
}
