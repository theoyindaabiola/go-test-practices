package main

import (
	"net/http" // allows hhtp requests
	"strconv"

	"github.com/gin-gonic/gin" // web framework
)

/**
	this is the logic side of things that puts in the actual task content from the user
	it works with the http request and webframe work to respond to the users requests
	Basically interaction logic
**/

/**
	gin is the web framework, the context represents request and response access the payload, URL, headers etc...
**/

func CreateTask(c *gin.Context) {
	var task Task

	// bind the JSON
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
/**
	doesn't need "return err" because gin gets us the error
	the payload coming from the user is the JSON right, gin converts the 
	response to the user to JSON,  and if any error, return the error in JSON format as well
**/
	if err := CreateTaskDB(task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, "Task Created")
}

func GetTasks(c *gin.Context) {
	// call the function 
	tasks, err := GetTasksDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func GetTask(c *gin.Context) {
	// needed as the key, coming from the URL request 
	id := c.Param("id")
	// call the function 
	task, err := GetTaskDB(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, task)
}

func UpdateTask(c *gin.Context) {
	// needed as the key, coming from the URL request
	id := c.Param("id")
	// convert id to string using strconv
	updateID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		// use gin's error
		c.JSON(400, gin.H{"error": "Invalid task ID."})
	}
	// need an interface placeholder for the properties to be upadted
	var task map[string]interface{}

	// bcos its a payload, it needs to be binded
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// update here
	if err := UpdateTaskDB(uint(updateID), task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	if err := DeleteTaskDB(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, "Task successfully deleted")
}
