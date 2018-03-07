package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gtongy/demo-echo-app/models"
	"github.com/labstack/echo"
)

func GetUsers(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, models.GetUsers(db))
	}
}
