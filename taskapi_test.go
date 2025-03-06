package main

import (
	"bytes"
	"encoding/json"
	"strconv"

	// "encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Connect to database using the sqlite
func setupTestDB() {
	var err error
	// uses the sqlite to use the local memorydisk machine to store data
	db, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		// panic is for test
		panic("Failed to connect to the database")
	}
	db.AutoMigrate(&Task{})
}

// Set up the router
func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	
	// group route
	// taskRoute := r.Group("/")
	r.POST("/tasks", CreateTask)
	r.GET("/tasks", func(c *gin.Context)  {
		// store an array of multiple task in the variable task
		var tasks []Task
		db.Find(&tasks)
		c.JSON(http.StatusOK, tasks)
	})
	r.PUT("/tasks/:id", UpdateTask)
	r.DELETE("/tasks/:id", DeleteTask)
	return r
}