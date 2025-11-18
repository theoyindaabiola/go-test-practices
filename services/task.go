package services

import (
	"taskapi/dao"
	"taskapi/models"
)

type TaskService struct {
	TaskDAO *dao.TaskDAO
}

func NewTaskService(s *dao.TaskDAO) *TaskService {
	return &TaskService{TaskDAO: s}
}

func (s *TaskService) CreateTask(task models.Task) error {
	return s.TaskDAO.CreateTaskDB(task)
}

func (s *TaskService) GetTasks() ([]models.Task, error) {
	return s.TaskDAO.GetTasksDB()
}

func (s *TaskService) GetTask(id string) (models.Task, error) {
	return s.TaskDAO.GetTaskDB(id)
}

func (s *TaskService) UpdateTask(id string, task map[string]interface{}) error {
	return s.TaskDAO.UpdateTaskDB(id, task)
}

func (s *TaskService) DeleteTask(id string) error {
	return s.TaskDAO.DeleteTaskDB(id)
}
