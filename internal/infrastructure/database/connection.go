package database

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	var DB *gorm.DB
	db_url := os.Getenv("DATABASE_URL")
	var err error
	DB, err = gorm.Open(postgres.Open(db_url), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}
	//DB.AutoMigrate(&domain.User{}, &domain.Activity{})

	return DB
}
