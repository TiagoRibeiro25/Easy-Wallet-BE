package main

import (
	"fmt"
	"net/http"

	"easy-wallet-be/src/configs"

	"github.com/labstack/echo/v4"
)

// The function creates an HTTP server using the Echo framework
func Server() *echo.Echo {
	fmt.Println("Starting the http server...")

	server := echo.New()

	server.GET("/api/v1", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Hello World!",
		})
	})

	return server
}

func main() {
	// Clear terminal
	fmt.Print("\033[H\033[2J")

	// Load envs from .env file
	configs.LoadEnv()

	// Validate required envs
	configs.ValidateEnvs()

	// Instantiate the server
	server := Server()

	server.Logger.Fatal(server.Start(fmt.Sprintf(":%s", configs.GetEnv("PORT"))))

	defer server.Close()
}
