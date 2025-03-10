package main

import (
	"bytes"
	"encoding/json"
	"strconv"
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
	r := gin.Default() // create a router for me
	
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

// functions unit testing
func TestTaskAPI(t *testing.T) {
	setupTestDB()
	router := setupRouter()

	// CREATE TASK
	taskData := `{"title": "Oyindamola, The Mastermind", "completed": "true"}`
	// request
	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(([]byte(taskData))))
	// header
	req.Header.Set("Content-Type", "application/json")
	// test endpoint
	w := httptest.NewRecorder()
	// route the test and req
	router.ServeHTTP(w, req)
	// assert 
	assert.Equal(t, http.StatusCreated, w.Code)

	// TEST GET TASK
	// GET request
	req, _ = http.NewRequest("GET", "/tasks", nil)
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	// confirm the response
	var tasks []Task // placeholder for the retrieved tasks
	// get the bodyconvert to what go understands and put in memory location of task
	json.Unmarshal(w.Body.Bytes(), &tasks)
	// assert the length retrieved
	assert.Len(t, tasks, 1)


	// TEST UPDATE TASK
	// get the ID of the tasked fetched above
	taskID := tasks[0].ID
	// create a tweaked data to be updated
	updatedData := `{"title": "Oyindamola's weekly report", "completed": "true"}`
	// use the PUT request to updated
	req, _ = http.NewRequest("PUT", "/tasks/"+strconv.Itoa(int(taskID)), bytes.NewBuffer(([]byte(updatedData))))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	// DELETE TASK 
	
}