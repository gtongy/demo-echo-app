package models

import (
	"fmt"
	"time"

	"github.com/gtongy/demo-echo-app/mysql"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uint   `gorm:"primary_key"`
	Email     string `validate:"required,email"`
	Password  string `validate:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) Auth(password string) error {
	hash := u.Password
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err
}

func PasswordHash(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}
	return string(hash)
}

func (user *User) Get() {
	db := mysql.GetDB()
	db.Find(&user)
}
