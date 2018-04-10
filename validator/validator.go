package validator

import (
	"net/http"

	"github.com/gtongy/demo-echo-app/errors"
	"github.com/gtongy/demo-echo-app/models"
	"github.com/gtongy/demo-echo-app/mysql"
	"github.com/labstack/echo"
	validator "gopkg.in/go-playground/validator.v9"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}

func New() *validator.Validate {
	return validator.New()
}

func ApiAccessTokenValidator(key string, c echo.Context) (bool, error) {
	var user models.User
	db := mysql.GetDB()
	defer db.Close()
	err := db.Where("access_token = ?", key).Find(&user).Error
	if err != nil {
		return false, errors.APIError.JSONErrorHandler(err, c, http.StatusUnauthorized, "Access token is invalid")
	}
	return true, nil
}
