package main

import (
	"html/template"
	"io"
	"os"

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
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	store := redis.GetStore()

	app := e.Group("")
	app.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "form:csrf",
	}))

	app.Use(session.Middleware(store))

	app.Static("/css", "./assets/css")

	app.GET("/", handlers.User.Top)
	app.GET("/login", handlers.User.Login)
	app.GET("/logout", handlers.User.Logout)
	app.GET("/register", handlers.User.Register)
	app.POST("/user/create", handlers.User.Create)
	app.POST("/auth", handlers.User.Auth)

	api := e.Group("/api/v1")
	api.Use(middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		KeyLookup: "header:DEMO-ECHO-TOKEN",
		Validator: validator.ApiAccessTokenValidator,
	}))

	api.GET("/tasks", handlers.Task.Get)
	api.POST("/tasks", handlers.Task.Create)

	e.Logger.Fatal(e.Start(":" + port()))
}

func port() string {
	defaultPort := "1323"
	envPort := os.Getenv("PORT")
	if envPort == "" {
		return defaultPort
	}
	return envPort
}
