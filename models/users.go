package models

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
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

func PasswordHash(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}
	return string(hash)
}
