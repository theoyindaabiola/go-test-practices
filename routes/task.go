package routes

import(
	"taskapi/controllers"
	"github.com/gin-gonic/gin"
)

type Routes struct {
	TaskController *controllers.TaskController
}

func RegisterRoutes(router *gin.Engine, taskController *controllers.TaskController) {
	// group the routes
	taskRoutes := router.Group("/api/tasks")
	{
		taskRoutes.POST("/", taskController.CreateTask)
		taskRoutes.GET("/", taskController.GetTasks)
		taskRoutes.GET("/:id", taskController.GetTask)
		taskRoutes.PUT("/:id", taskController.UpdateTask)
		taskRoutes.DELETE("/:id", taskController.DeleteTask)
	}
}