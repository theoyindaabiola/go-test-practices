package main

// the struct is used in both JSON and database, thus it is different from the regular struct
type Task struct {
	ID 			uint `gorm:"primaryKey" json:"id"`
	Title 		string `gorm:"not null" json:"title"`
	Completed 	string `gorm:"default:false" json:"completed"`
}
