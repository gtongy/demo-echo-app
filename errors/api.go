package errors

import (
	"github.com/labstack/echo"
)

var APIError apiError

type apiError struct {
	Code    int
	Message string
}

func (apiError *apiError) JSONErrorHandler(err error, c echo.Context, code int, message string) error {
	apiError.Code = code
	apiError.Message = message
	c.JSON(code, apiError)
	return err
}
