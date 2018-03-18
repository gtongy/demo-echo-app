package models

import (
	"database/sql"
)

// User is a struct containing User data
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

// UserCollection is collection of Users
type UserCollection struct {
	Users []User `json:"items"`
}

// GetUsers from the DB
func GetUsers(db *sql.DB) UserCollection {
	sql := "SELECT * FROM users"
	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	result := UserCollection{}
	for rows.Next() {
		User := User{}
		scanErr := rows.Scan(&User.ID, &User.Name, &User.Password)
		if scanErr != nil {
			panic(scanErr)
		}
		result.Users = append(result.Users, User)
	}
	return result
}
