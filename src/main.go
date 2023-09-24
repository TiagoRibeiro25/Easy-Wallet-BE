package main

import (
	"fmt"

	"easy-wallet-be/src/models"
	"easy-wallet-be/src/utils"
)

// The main function initializes and starts a server, loading environment variables and validating
// required variables before starting.
func main() {
	// Clear terminal
	fmt.Print("\033[H\033[2J")

	// Connect to the database
	db := models.DB()

	// Instantiate the server
	server := Server()
	server.Logger.Fatal(server.Start(fmt.Sprintf(":%s", utils.GetEnv("PORT"))))

	// Defer closing the database connection and server
	defer func() {
		db.Close()
		server.Close()
	}()
}
