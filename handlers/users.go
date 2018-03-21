package handlers

import (
	"net/http"

	"github.com/gtongy/demo-echo-app/models"
	"github.com/gtongy/demo-echo-app/mysql"
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
	password := models.PasswordHash(c.FormValue("password"))
	user := models.User{
		Email:    email,
		Password: password,
	}
	if err := c.Bind(user); err != nil {
		return err
	}

	db := mysql.GetDB()
	defer db.Close()
	db.Create(&user)
	return c.Redirect(http.StatusMovedPermanently, "/login")
}

func (u *user) Auth(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")
	user := models.User{
		Email:    email,
		Password: password,
	}
	db := mysql.GetDB()
	defer db.Close()
	db.Where("email = ?", user.Email).First(&user)
	err := user.Auth(password)
	if err != nil {
		return c.Redirect(http.StatusMovedPermanently, "/login")
	}
	// TODO: make home template and replace redirect
	return c.Redirect(http.StatusMovedPermanently, "/register")
}
