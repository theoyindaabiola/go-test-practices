package main

/**
	log is used for logging detailed behaviour in terminal while fmt is for error handling
**/

import (
	"log"
	"taskapi/routes"
	"taskapi/dao"
	"taskapi/services"
	"taskapi/controllers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// load the .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("failed to load .env file: %w", err)
	}

	// connect to the database
	db := ConnectDB()

	// use gin to handle the routes
	router := gin.Default()

	// initialize the dao, service and controller
	taskDao := dao.NewTaskDAO(db)
	taskService := services.NewTaskService(taskDao)
	taskController := controllers.NewTaskController(taskService)

	// register the routes
	routes.RegisterRoutes(router, taskController)
	

	// starts the server
	log.Println("Server started at: 8080...")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("failed to start the server: %v", err)
	}
}