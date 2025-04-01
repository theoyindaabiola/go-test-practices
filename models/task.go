package models

type Task struct {
    ID          uint   `gorm:"primaryKey" json:"id"`
    Title       string `gorm:"not null" json:"title"`
    Description string `gorm:"not null" json:"description"`
    Completed   string `gorm:"default:false" json:"completed"`
}