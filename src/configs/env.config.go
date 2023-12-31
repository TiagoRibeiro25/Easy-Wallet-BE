package configs

import (
	"encoding/json"
	"fmt"
	"os"

	"easy-wallet-be/src/utils"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

func init() {
	// Load envs from .env file
	if !utils.IsProduction() {
		loadEnvs()
	}

	// Validate required envs
	validateEnvs()
}

// `loadEnvs` loads environment variables from a .env file using the godotenv package.
// If no path is provided, it defaults to ".env" in the current working directory.
func loadEnvs(paths ...string) {
	color.Cyan("Loading environment variables...")

	var path string

	// Check if a path is provided, otherwise use the default ".env" file in the current working directory.
	if len(paths) > 0 {
		path = paths[0]
	} else {
		path = ".env"
	}

	err := godotenv.Load(path)
	utils.HandleError(err, "Error loading .env file", true)

	color.Green("Environment variables loaded successfully")
}

// The function `validateEnvs` reads a JSON file containing a list of environment variable names,
// checks if each variable is set, and panics if any variable is not set.
func validateEnvs() {
	color.Cyan("Validating environment variables...")

	// Open the JSON file containing the list of environment variables
	file, err := os.Open("src/data/envs.json")
	utils.HandleError(err, "", true)

	defer file.Close()

	var envs []string

	// Decode the JSON file into a slice of strings
	decoder := json.NewDecoder(file)
	utils.HandleError(decoder.Decode(&envs), "", true)

	// Check if each environment variable is set
	for _, env := range envs {
		if os.Getenv(env) == "" {
			utils.HandleError(fmt.Errorf("environment variable %s is not set", env), "", true)
		}
	}

	color.Green("All required environment variables are set")
}
