package dao

import (
	"gorm.io/gorm"
	"taskapi/models"
)

type TaskDAO struct {
	DB *gorm.DB
}

func NewTaskDAO(db *gorm.DB) *TaskDAO {
	return &TaskDAO{DB: db}
}

func (dao *TaskDAO) CreateTaskDB(task models.Task) error {
	if err := dao.DB.Create(&task).Error; err != nil {
		return err
	}

	return nil
}

func (dao *TaskDAO) GetTasksDB() ([]models.Task, error) {
	var tasks []models.Task
	// get all selected from the db and pore them into tasks's memory location
	if err := dao.DB.Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (dao *TaskDAO) GetTaskDB(id uint) (models.Task, error) {
	var task models.Task
	// get all selected from the db and pore them into tasks's memory location
	if err := dao.DB.Where("id = ?", id).First(&task).Error; err != nil {
		return task, err
	}
	return task, nil
}

// here the GORM accepts map as struct for updating, empty interface is flexible for updating, struct is not.
func (dao *TaskDAO) UpdateTaskDB(id uint, task map[string]interface{}) error {
	// placeholder for the task to be updated
	var updateTask models.Task
	// finds the task by id and store in the memory location of updateTask
	if err := dao.DB.Where("id = ?", id).First(&updateTask).Error; err != nil {
		return err
	}
	// return the fetched task to be updated and update the interface values of task
	return dao.DB.Model(&updateTask).Updates(task).Error
}

func (dao *TaskDAO) DeleteTaskDB(id uint) error {
	var task models.Task
	// get all selected from the db and pore them into tasks's memory location
	if err := dao.DB.Where("id = ?", id).First(&task).Error; err != nil {
		return err
	}
	// gorm functions have Error
	return dao.DB.Delete(&task).Error
}
