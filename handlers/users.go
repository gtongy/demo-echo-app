package handlers

import (
	"net/http"

	"github.com/gtongy/demo-echo-app/models"
	"github.com/gtongy/demo-echo-app/mysql"
	"github.com/gtongy/demo-echo-app/redis"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
)

var User user

type user struct{}

func (u *user) Top(c echo.Context) error {
	sess, _ := session.Get("session", c)
	// TODO: change type pushy.
	user := models.User{
		ID: sess.Values["userId"].(uint),
	}
	user.Get()
	return c.Render(http.StatusOK, "content", map[string]interface{}{
		"user": user,
	})
}

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
	session := redis.GetSession(c)
	session.Values["userId"] = user.ID
	session.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusMovedPermanently, "/")
}
