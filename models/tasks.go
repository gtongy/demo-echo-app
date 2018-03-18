package models

import (
	"database/sql"
)

// Task is a struct containing Task data
type Task struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// TaskCollection is collection of Tasks
type TaskCollection struct {
	Tasks []Task `json:"items"`
}

// GetTasks from the DB
func GetTasks(db *sql.DB) TaskCollection {
	sql := "SELECT * FROM tasks"
	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	result := TaskCollection{}
	for rows.Next() {
		task := Task{}
		scanErr := rows.Scan(&task.ID, &task.Name)
		if scanErr != nil {
			panic(scanErr)
		}
		result.Tasks = append(result.Tasks, task)
	}
	return result
}
