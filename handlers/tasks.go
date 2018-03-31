package handlers

import (
	"net/http"
	"strconv"

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
	return c.JSON(http.StatusOK, task)
}

func (t *task) Create(c echo.Context) error {
	title := c.FormValue("title")
	userID, err := strconv.Atoi(c.FormValue("user_id"))
	if err != nil {
		return err
	}
	task := &models.Task{
		Title:  title,
		UserID: userID,
	}
	if err := c.Bind(task); err != nil {
		return err
	}
	db := mysql.GetDB()
	defer db.Close()
	db.Create(&task)
	return c.JSON(http.StatusOK, task)
}
