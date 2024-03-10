package pg

import (
	"fmt"

	"brainhabit/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(dsn string) (*gorm.DB, error) {

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("Failed to connect to postgres: %w", err)
	}

	fmt.Println("Server connected to postgres")

	err = AutoMigrate(db)

	if err != nil {
		return nil, fmt.Errorf("Failed automigrate postgres database: %w", err)
	}

	fmt.Println("Postgres automigration successful")

	return db, nil
}

func AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(&models.User{})
	if err != nil {
		return err
	}
	return nil
}
