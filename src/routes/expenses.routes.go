package routes

import (
	handlers "easy-wallet-be/src/handlers/expenses"
	"easy-wallet-be/src/middlewares"

	"github.com/labstack/echo/v4"
)

// ExpensesRoutes sets up the expenses routes for the server.
func ExpensesRoutes(server *echo.Group) {
	expensesRoutes := server.Group("/v1/expenses")

	// Apply session validation middleware to all routes
	expensesRoutes.Use(middlewares.ValidateSessionMiddleware)

	// Get expenses
	expensesRoutes.GET("", handlers.GetExpenses)

	// Add a new expense
	expensesRoutes.POST(
		"",
		handlers.AddExpense,
		middlewares.ValidateJSONSchema("expenses/addExpense"),
	)
}
