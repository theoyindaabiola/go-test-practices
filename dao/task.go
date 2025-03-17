package dao

import (
	"gorm.io/gorm"
)

func CreateTaskDB(task Task) error {
	if err := db.Create(&task).Error; err != nil {
		return err
	}

	return nil
}

func GetTasksDB() ([]Task, error) {
	var tasks []Task
	// get all selected from the db and pore them into tasks's memory location
	if err := db.Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func GetTaskDB(id uint) (Task, error) {
	var task Task
	// get all selected from the db and pore them into tasks's memory location
	if err := db.Where("id = ?", id).First(&task).Error; err != nil {
		return task, err
	}
	return task, nil
}

// here the GORM accepts map as struct for updating, empty interface is flexible for updating, struct is not.
func UpdateTaskDB(id uint, task map[string]interface{}) error {
	// placeholder for the task to be updated
	var updateTask Task
	// finds the task by id and store in the memory location of updateTask
	if err := db.Where("id = ?", id).First(&updateTask).Error; err != nil {
		return err
	}
	// return the fetched task to be updated and update the interface values of task
	return db.Model(&updateTask).Updates(task).Error
}

func DeleteTaskDB(id uint) error {
	var task Task
	// get all selected from the db and pore them into tasks's memory location
	if err := db.Where("id = ?", id).First(&task).Error; err != nil {
		return err
	}
	// gorm functions have Error
	return db.Delete(&task).Error
}

/// function testing with air...
/// unit testing...