package main

// the struct is used in both JSON and database, thus it is different from the regular struct
type Task struct {
	ID string `json:"id" db:"id"`
	Title string `json:"title" db:"title"`
	Completed string `json:"completed" db:"completed"`
}
