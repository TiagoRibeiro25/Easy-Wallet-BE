package configs

import "os"

type DatabaseConfig struct {
	Host      string
	Port      string
	Username  string
	Password  string
	Database  string
	EnableSSL bool
}

// The function GetDatabaseConfig retrieves the database configuration from environment variables.
func GetDatabaseConfig() DatabaseConfig {

	return DatabaseConfig{
		Host:      os.Getenv("DB_HOST"),
		Port:      os.Getenv("DB_PORT"),
		Username:  os.Getenv("DB_USERNAME"),
		Password:  os.Getenv("DB_PASSWORD"),
		Database:  os.Getenv("DB_NAME"),
		EnableSSL: os.Getenv("DB_ENABLE_SSL") == "true",
	}
}
