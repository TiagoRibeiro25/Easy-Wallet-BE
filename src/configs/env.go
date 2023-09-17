package configs

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv loads environment variables from a .env file using the godotenv package.
// If no path is provided, it defaults to ".env" in the current working directory.
func LoadEnv(paths ...string) {
	fmt.Println("Loading environment variables...")

	var path string

	// Check if a path is provided, otherwise use the default ".env" file in the current working directory.
	if len(paths) > 0 {
		path = paths[0]
	} else {
		path = ".env"
	}

	err := godotenv.Load(path)

	if err != nil {
		panic("Error loading .env file")
	}

	fmt.Println("Environment variables loaded successfully")
}

// The function GetEnv retrieves the value of an environment variable given its key.
func GetEnv(key string) string {
	return os.Getenv(key)
}

// The function `ValidateEnvs` reads a JSON file containing a list of environment variable names,
// checks if each variable is set, and panics if any variable is not set.
func ValidateEnvs() {
	fmt.Println("Validating environment variables...")

	// Open the JSON file containing the list of environment variables
	file, err := os.Open("src/data/envs.json")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	var envs []string

	// Decode the JSON file into a slice of strings
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&envs); err != nil {
		panic(err)
	}

	// Check if each environment variable is set
	for _, env := range envs {
		if GetEnv(env) == "" {
			panic(fmt.Sprintf("Environment variable %s is not set", env))
		}
	}

	fmt.Println("All required environment variables are set")
}
