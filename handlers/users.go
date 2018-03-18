package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

func Login(c echo.Context) error {
	return c.Render(http.StatusOK, "content", map[string]interface{}{
		"name": "Dolly!",
	})
}
