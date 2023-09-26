package models

import (
	"easy-wallet-be/src/configs"
	"easy-wallet-be/src/utils"

	"github.com/fatih/color"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB // Package-level database connection variable

// Initialize the database connection
func init() {
	var err error
	db, err = SetupDatabase()
	utils.HandleError(err, "", true)

	// Automatically create the database tables
	// setUpModels()
}

// The function `SetupDatabase` sets up a connection to a database using the provided configuration.
func SetupDatabase() (*gorm.DB, error) {
	color.Cyan("Connecting to the database...")

	dbConfig := configs.GetDatabaseConfig()
	var ssl string

	if dbConfig.EnableSSL {
		ssl = "require"
	} else {
		ssl = "disable"
	}

	dsn := "postgres://" +
		dbConfig.Username + ":" +
		dbConfig.Password + "@" +
		dbConfig.Host + ":" +
		dbConfig.Port + "/" +
		dbConfig.Database + "?sslmode=" +
		ssl

	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	color.Green("Successfully connected to the database")

	return db, nil
}

// `setUpModels` sets up the database models by creating the necessary tables and relationships.
// It also retrieves the first user and password from the database and preloads their associated models.
func setUpModels() {
	color.Cyan("Setting up models...")

	// AutoMigrate the models to create the necessary tables
	err := db.AutoMigrate(
		&User{},
		&Password{},
		&Session{},
		&Income{},
		&Expense{},
		&Category{},
	).Error
	utils.HandleError(err, "Error migrating models", true)

	color.Green("Successfully set up models")
}

// `DB` returns the shared database connection
func DB() *gorm.DB {
	return db
}
