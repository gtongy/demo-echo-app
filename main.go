package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"io"

	"github.com/gtongy/demo-echo-app/handlers"

	"github.com/BurntSushi/toml"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Config struct {
	DataSource string `toml:"dataSource"`
}

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()
	var config Config
	_, err := toml.DecodeFile("./config.toml", &config)
	if err != nil {
		panic(err)
	}
	dataSource := config.DataSource
	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		fmt.Println(err.Error())
	}

	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("templates/*.tpl")),
	}
	e.Renderer = renderer

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/tasks", handlers.GetTasks(db))
	e.GET("/users", handlers.GetUsers(db))
	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "content", map[string]interface{}{
			"name": "Dolly!",
		})
	})

	e.Logger.Fatal(e.Start(":1323"))
}
