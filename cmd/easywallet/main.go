package main

import (
	"fmt"

	"easy-wallet-be/internal/config"
)

func main() {
	config.LoadEnv()
	config.ValidateEnvs()

	fmt.Println("Hello, World!")
}
