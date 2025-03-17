package main

/** 
	log is used for logging detailed behaviour in terminal while fmt is for error handling
**/

import(
	"log" 
	"github.com/gin-gonic/gin"
)

func main() {
	// connect to the database
	ConnectDB()

	// use gin to handle the routes
	router := gin.Default()

	router.POST("/tasks", CreateTask)
	router.GET("/tasks", GetTasks)
	router.GET("/tasks/:id", GetTask)
	router.PUT("/tasks/:id", UpdateTask)
	router.DELETE("/tasks/:id", DeleteTask)

	// starts the server
	log.Println("Server started at: 8080...")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("failed to start the server: %v", err)
	}
}