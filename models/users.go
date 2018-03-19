package models

import (
	"database/sql"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserCollection struct {
	users []User `json:"items"`
}

// Getusers from the DB
func (u *User) Create(db *sql.DB) UserCollection {

	sql := "SELECT * FROM users"
	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	result := UserCollection{}
	for rows.Next() {
		user := User{}
		scanErr := rows.Scan(&user.Email, &user.Password)
		if scanErr != nil {
			panic(scanErr)
		}
		result.users = append(result.users, user)
	}
	return result
}
