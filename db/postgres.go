package pg

import (
	"fmt"

	"brainhabit/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Setup(dsn string) error {

	db, err := Connect(dsn)

	if err != nil {
		return err
	}

	err = AutoMigrate(db)

	if err != nil {
		return fmt.Errorf("Error, failed automigrate postgres schema: %w", err)
	}

	DB = db

	return nil
}

func Connect(dsn string) (*gorm.DB, error) {

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("Error, failed to connect to postgres: %w", err)
	}

	fmt.Println("Server connected to postgres")

	return database, nil

}

func AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(&models.User{})
	if err != nil {
		return err
	}
	fmt.Println("Postgres schema automigration successful")
	return nil
}
