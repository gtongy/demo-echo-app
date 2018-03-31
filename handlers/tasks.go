package handlers

import (
	"net/http"

	"github.com/gtongy/demo-echo-app/models"
	"github.com/gtongy/demo-echo-app/mysql"
	"github.com/labstack/echo"
)

var Task task

type task struct{}

func (t *task) Get(c echo.Context) error {
	var task []models.Task
	db := mysql.GetDB()
	db.Find(&task)
	return c.JSON(http.StatusOK, task)
}
