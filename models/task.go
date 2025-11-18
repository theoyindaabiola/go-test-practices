package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Task struct {
    ID              string      `gorm:"primaryKey" json:"id"`
    Title           string      `gorm:"not null" json:"title"`
    Description     string      `gorm:"not null" json:"description"`
    Completed       bool        `gorm:"default:false" json:"completed"`
}


func (t *Task) BeforeCreate(tx *gorm.DB) error {
	if t.ID == "" {
		t.ID = uuid.NewString()
	}	
	return nil
}
