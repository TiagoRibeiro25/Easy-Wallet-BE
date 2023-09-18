package main

import (
	"fmt"

	"easy-wallet-be/src/configs"
	"easy-wallet-be/src/utils"
)

func main() {
	// Clear terminal
	fmt.Print("\033[H\033[2J")

	// Load envs from .env file
	configs.LoadEnv()

	// Validate required envs
	configs.ValidateEnvs()

	// Instantiate the server
	server := Server()

	server.Logger.Fatal(server.Start(fmt.Sprintf(":%s", utils.GetEnv("PORT"))))

	defer server.Close()
}
