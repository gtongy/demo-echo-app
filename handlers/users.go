package handlers

import (
	"fmt"
	"net/http"

	"github.com/gtongy/demo-echo-app/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
)

var User user

type user struct{}

func (u *user) Login(c echo.Context) error {
	return c.Render(http.StatusOK, "form", map[string]interface{}{})
}

func (u *user) Register(c echo.Context) error {
	return c.Render(http.StatusOK, "form", map[string]interface{}{
		"new": true,
	})
}

func (u *user) Create(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")
	user := &models.User{
		Email:    email,
		Password: password,
	}
	var err = c.Bind(user)
	if err != nil {
		fmt.Println(err)
	}
	db, err := gorm.Open("mysql", "root:root@tcp(mysql:3306)/test_database")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.Create(&user)
	return c.Redirect(http.StatusMovedPermanently, "/login")
}
