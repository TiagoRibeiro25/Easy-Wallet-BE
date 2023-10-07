package routes

import (
	handlers "easy-wallet-be/src/handlers/categories"
	"easy-wallet-be/src/middlewares"

	"github.com/labstack/echo/v4"
)

// CategoriesRoutes sets up the categories routes for the server.
func CategoriesRoutes(server *echo.Group) {
	categoriesRoutes := server.Group("/v1/categories")

	// Apply session validation middleware to all routes
	categoriesRoutes.Use(middlewares.ValidateSessionMiddleware)

	// Get all user categories
	categoriesRoutes.GET("", handlers.GetCategories)

	// Add a new category
	categoriesRoutes.POST(
		"",
		handlers.AddCategory,
		middlewares.ValidateJSONSchema("categories/addCategory"),
	)

	// Delete a category
	categoriesRoutes.DELETE("/:id", handlers.DeleteCategory)
}
