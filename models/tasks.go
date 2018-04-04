package models

// Task is a struct containing Task data
type Task struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Title  string `json:"title"`
	UserID uint   `json:"user_id"`
}
