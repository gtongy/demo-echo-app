package handlers

import (
	"net/http"

	"github.com/gtongy/demo-echo-app/errors"
	"github.com/gtongy/demo-echo-app/models"
	"github.com/gtongy/demo-echo-app/mysql"
	"github.com/labstack/echo"
)

var Task task

type task struct{}

func (t *task) Get(c echo.Context) error {
	var task []models.Task
	db := mysql.GetDB()
	defer db.Close()
	db.Find(&task)
	tasks := models.Tasks{Tasks: task}
	return c.JSON(http.StatusOK, tasks)
}

func (t *task) Create(c echo.Context) error {
	var user models.User
	db := mysql.GetDB()
	defer db.Close()
	db.Where("access_token = ?", c.Request().Header.Get("DEMO-ECHO-TOKEN")).Find(&user)
	title := c.FormValue("title")

	task := &models.Task{
		Title:  title,
		UserID: user.ID,
	}

	if err := c.Bind(task); err != nil {
		return errors.APIError.JSONErrorHandler(err, c, http.StatusBadRequest, "Request is invalid")
	}

	if err := c.Validate(task); err != nil {
		return errors.APIError.JSONErrorHandler(err, c, http.StatusBadRequest, "Validate is failed")
	}

	db.Create(&task)
	return c.JSON(http.StatusOK, task)
}
