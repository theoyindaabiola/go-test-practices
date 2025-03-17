package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

/**
	This is the database creation side, focus on the table and CRUD operations
	it works with the http request and webframe work *
**/

var db *gorm.DB

func ConnectDB() {

	var err error

	if err := godotenv.Load(); err != nil {
		fmt.Printf("failed to load .env file: %v", err)
	}

	environmentStr := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)

	db, err = gorm.Open(postgres.Open(environmentStr), &gorm.Config{})
	if err != nil {
		fmt.Printf("failed to connect to database: %v\n", err)
	}

	db.AutoMigrate(&Task{});
}
