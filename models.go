package main

// the struct is used in both JSON and database, thus it is different from the regular struct
type Task struct {
	ID string `gorm:"primaryKey"`
	Title string `gorm:"not null"`
	Completed string `gorm:"default:false"`
}
