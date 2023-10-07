package routes

import (
	handlers "easy-wallet-be/src/handlers/auth"
	"easy-wallet-be/src/middlewares"

	"github.com/labstack/echo/v4"
)

// AuthRoutes sets up the authentication routes for the server.
func AuthRoutes(server *echo.Group) {
	authRoutes := server.Group("/auth")

	// Register new user
	authRoutes.POST(
		"/register",
		handlers.Register,
		middlewares.ValidateJSONSchema("auth/register"),
	)

	// Login user
	authRoutes.POST(
		"/login",
		handlers.Login,
		middlewares.ValidateJSONSchema("auth/login"),
	)

	// Logout user
	authRoutes.DELETE("/logout", handlers.Logout)

	// Forgot password (send reset password email)
	authRoutes.POST(
		"/forgot-password",
		handlers.ForgotPassword,
		middlewares.ValidateJSONSchema("auth/forgotPassword"),
	)

	// Reset password
	authRoutes.PATCH(
		"/reset-password/:token",
		handlers.ResetPassword,
		middlewares.ValidateJSONSchema("auth/resetPassword"),
	)
}
