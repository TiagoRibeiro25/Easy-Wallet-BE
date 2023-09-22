package configs

import "easy-wallet-be/src/utils"

type DatabaseConfig struct {
	Dialect   string
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
		Dialect:   utils.GetEnv("DB_DIALECT"),
		Host:      utils.GetEnv("DB_HOST"),
		Port:      utils.GetEnv("DB_PORT"),
		Username:  utils.GetEnv("DB_USERNAME"),
		Password:  utils.GetEnv("DB_PASSWORD"),
		Database:  utils.GetEnv("DB_NAME"),
		EnableSSL: utils.GetEnv("DB_ENABLE_SSL") == "true",
	}
}
