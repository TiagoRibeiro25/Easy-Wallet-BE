package models

import (
	"easy-wallet-be/src/configs"

	"github.com/fatih/color"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// The function `SetupDatabase` sets up a connection to a database using the provided configuration.
func SetupDatabase() (*gorm.DB, error) {
	color.Cyan("Connecting to the database...")

	dbConfig := configs.GetDatabaseConfig()
	var ssl string

	if dbConfig.EnableSSL {
		ssl = "enable"
	} else {
		ssl = "disable"
	}

	dsn := dbConfig.Dialect + "://" +
		dbConfig.Username + ":" +
		dbConfig.Password + "@" +
		dbConfig.Host + ":" +
		dbConfig.Port + "/" +
		dbConfig.Database + "?sslmode=" +
		ssl

	db, err := gorm.Open(dbConfig.Dialect, dsn)
	if err != nil {
		return nil, err
	}

	color.Green("Successfully connected to the database")

	return db, nil
}
