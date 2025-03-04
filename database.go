package main

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_"github.com/lib/pq"
)

/** 
	This is the database creation side, focus on the table and CRUD operations
	it works with the http request and webframe work *
**/

var db *sqlx.DB

func ConnectDB() error {
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("failed to load .env file: %w", err)
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

	var err error

	db, err = sqlx.Connect("postgres", environmentStr)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	if err := AutoMigrate(); err != nil {
		return fmt.Errorf("failed to migrate: %v", err)
	}
	return nil
}

func AutoMigrate() error {
	query := `
		CREATE TABLE IF NOT EXISTS tasks (
			id VARCHAR(36) PRIMARY KEY,
			title TEXT NOT NULL,
			completed BOOLEAN DEFAULT FALSE
		);
	`

	_, err := db.Exec(query)
	return err
}

func CreateTaskDB(task Task) error {
	query := `INSERT INTO tasks (id, title, completed) VALUES ($1, $2, $3)`
	_, err := db.Exec(query, task.ID, task.Title, task.Completed)

	return err
}

func GetTasksDB() ([]Task, error) {
	var tasks []Task
	query := `SELECT * FROM tasks`
	// get all selected from the query and pore them into tasks's memory
	err := db.Select(&tasks, query)
	return tasks, err
}

// 
// func GetTasksDB() ([]Task, error) {
//     var tasks []Task
//     query := `SELECT * FROM tasks`

//     if err := db.Get(&tasks, query); err != nil {
//         return tasks, err
//     }

//     return tasks, nil
// }

func GetTaskDB(id string) (Task, error) {
	var task Task
	query := `SELECT * FROM tasks WHERE id = $1`
	// get all selected from the query and pore them into tasks's memory
	err := db.Get(&task, query, id)
	return task, err
}

// func GetTaskDB(id string) (Task, error) {
// 	var task Task
// 	query := `SELECT * FROM tasks WHERE id = $1`
// 	if err := db.Get(&task, query); err != nil {
// 		return task, err
// 	}

// 	return task, nil
// }

func UpdateTaskDB(task Task) error {
	// query := `UPDATE tasks SET title = $1, completed = $2 id = $3`
	// _, err := db.Exec(query, task.ID, task.Title, task.Completed)

	query := `UPDATE tasks SET title = $1, completed = $2 WHERE id = $3`
	_, err := db.Exec(query, task.Title, task.Completed, task.ID)


	return err
}

/** 
	Exec returns 2 parameters hence "_, err" 
	Delete doesn't need a task placeholder because ...
	we're only picking the id and deleting it  

**/

func DeleteTaskDB(id string) error {
	query := `DELETE FROM tasks WHERE id = $1`
	_, err := db.Exec(query, id)

	return err
}

