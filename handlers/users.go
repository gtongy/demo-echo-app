package handlers

import (
	"fmt"
	"net/http"

	"github.com/gtongy/demo-echo-app/models"
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
	user := models.User{
		Email:    email,
		Password: password,
	}
	var err = c.Bind(user)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return c.JSON(http.StatusOK, user.Create)
}
