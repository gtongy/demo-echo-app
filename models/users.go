package models

import (
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	ID        uint `gorm:"primary_key"`
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserCollection struct {
	users []User
}
