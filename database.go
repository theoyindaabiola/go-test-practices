package main

import (
	"fmt"
	"os"
	"log"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

/**
	This is the database creation side, focus on the table creation and the connection to the database
**/

func ConnectDB() *gorm.DB {

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

	db, err := gorm.Open(postgres.Open(environmentStr), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database: %v\n", err)
	}

	err = db.AutoMigrate(&Task{});
	if err != nil {
		log.Fatal("failed to migrate the database: %v\n", err)
	}
	return db
}
