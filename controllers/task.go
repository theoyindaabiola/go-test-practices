package controllers

import (
	"net/http" // allows hhtp requests
	"taskapi/services"
	"taskapi/models" // the model

	"github.com/gin-gonic/gin" // web framework
)

type TaskController struct {
	TaskService *services.TaskService
}

func NewTaskController(tc *services.TaskService) *TaskController{
	return &TaskController{TaskService: tc}
}

/**
	this is the logic side of things that puts in the actual task content from the user
	it works with the http request and webframe work to respond to the users requests
	Basically interaction logic
**/

/**
	gin is the web framework, the context represents request and response access the payload, URL, headers etc...
**/

func (tc *TaskController) CreateTask(c *gin.Context) {
	var task models.Task

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
	if err := tc.TaskService.CreateTask(task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, "Task Created, successfully")
}

func (tc *TaskController) GetTasks(c *gin.Context) {
	// call the function 
	tasks, err := tc.TaskService.GetTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func (tc *TaskController) GetTask(c *gin.Context) {
	// needed as the key, coming from the URL request 
	id := c.Param("id")

	// // convert id to an uint
	// taskId, err := strconv.ParseUint(id, 10, 32)
	// if err != nil {
	// 	c.JSON(400, gin.H{"error": "Invalid task ID."})
	// }

	// call the function 
	task, err := tc.TaskService.GetTask(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, task)
}

func (tc *TaskController) UpdateTask(c *gin.Context) {
	id := c.Param("id") // needed as the key, coming from the URL request

	// // convert id to an uint
	// taskId, err := strconv.ParseUint(id, 10, 32)
	// if err != nil {
	// 	c.JSON(400, gin.H{"error": "Invalid task ID."})
	// }

	var task map[string]interface{}
	// get and confirm that there is no error with the payload
	if err := c.ShouldBindJSON(&task); err != nil {
		// if error, return error using http in JSON format using the gin context
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) 
		return
	}

	// task.ID = uint(taskId)
	if err := tc.TaskService.UpdateTask(id, task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, task)
}

func (tc *TaskController) DeleteTask(c *gin.Context) {
	id := c.Param("id")

	// // convert id to an uint
	// taskId, err := strconv.ParseUint(id, 10, 32)
	// if err != nil {
	// 	c.JSON(400, gin.H{"error": "Invalid task ID."})
	// }

	if err := tc.TaskService.DeleteTask(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// c.JSON(http.StatusNoContent, "Task successfully deleted")
	c.JSON(http.StatusCreated, "Task successfully deleted")
}