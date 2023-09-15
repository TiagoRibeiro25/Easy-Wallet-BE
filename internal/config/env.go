package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// The function `LoadEnv` loads environment variables from a .env file using the godotenv package
func LoadEnv() {
	err := godotenv.Load()

	if err != nil {
		panic("Error loading .env file")
	}
}

// The function GetEnv retrieves the value of an environment variable given its key.
func GetEnv(key string) string {
	return os.Getenv(key)
}

// The function `ValidateEnvs` reads a JSON file containing a list of environment variable names,
// checks if each variable is set, and panics if any variable is not set.
func ValidateEnvs() {
	file, err := os.Open("internal/data/envs.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var envs []string

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&envs); err != nil {
		panic(err)
	}

	for _, env := range envs {
		if GetEnv(env) == "" {
			panic(fmt.Sprintf("Environment variable %s is not set", env))
		}
	}
}
