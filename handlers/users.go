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
	password := c.FormValue("password")

	user := &models.User{
		Email:    email,
		Password: password,
	}

	var err = c.Bind(user)
	if err != nil {
		return err
	}

	db := mysql.GetDB()
	defer db.Close()
	db.Create(&user)
	return c.Redirect(http.StatusMovedPermanently, "/login")
}
