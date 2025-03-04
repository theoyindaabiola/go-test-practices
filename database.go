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
		fmt.Errorf("failed to load .env file: %w", err)
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
		fmt.Errorf("failed to connect to database: %w", err)
	}

	db.AutoMigrate(&Task{});
}

func CreateTaskDB(task Task) error {
	if err := db.Create(&task).Error; err != nil {
		return err
	}

	return nil
}

func GetTasksDB() ([]Task, error) {
	var tasks []Task
	// get all selected from the db and pore them into tasks's memory location
	if err := db.Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func GetTaskDB(id string) (Task, error) {
	var task Task
	// get all selected from the db and pore them into tasks's memory location
	if err := db.Where("id = ?", id).First(&task, id).Error; err != nil {
		return task, err
	}
	return task, nil
}

// here the GORM accepts map as struct for updating, empty interface is flexible for updating, struct is not.
func UpdateTaskDB(id uint, task map[string]interface{}) error {
	// placeholder for the task to be updated
	var updateTask Task
	// finds the task by id and store in the memory location of updateTask
	if err := db.Where("id = ?", id).First(&updateTask).Error; err != nil {
		return err
	}
	// return the fetched task to be updated and update the interface values of task
	return db.Model(&updateTask).Updates(task).Error
}

func DeleteTaskDB(id string) error {
	var task Task
	// get all selected from the db and pore them into tasks's memory location
	if err := db.Where("id = ?", id).First(&task, id).Error; err != nil {
		return err
	}
	// gorm functions have Error
	return db.Delete(&task).Error
}

