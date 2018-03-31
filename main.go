package main

import (
	"html/template"
	"io"

	"github.com/gtongy/demo-echo-app/handlers"
	"github.com/gtongy/demo-echo-app/redis"
	"github.com/gtongy/demo-echo-app/validator"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/middleware"
)

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()

	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("templates/*.tpl")),
	}
	e.Renderer = renderer
	e.Validator = &validator.CustomValidator{Validator: validator.New()}
	store := redis.Init()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "form:csrf",
	}))

	e.Use(session.Middleware(store))

	e.Static("/css", "./assets/css")

	e.GET("/", handlers.User.Top)
	e.GET("/login", handlers.User.Login)
	e.GET("/logout", handlers.User.Logout)
	e.GET("/register", handlers.User.Register)
	e.POST("/user/create", handlers.User.Create)
	e.POST("/auth", handlers.User.Auth)

	e.GET("/v1/tasks/", handlers.Task.Get)

	e.Logger.Fatal(e.Start(":1323"))
}
