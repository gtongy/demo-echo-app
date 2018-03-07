package main

import (
	"database/sql"
	"fmt"

	"github.com/BurntSushi/toml"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gtongy/demo-echo-app/handlers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Config struct {
	DataSource string `toml:"dataSource"`
}

func main() {
	// Echo instance
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
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/tasks", handlers.GetTasks(db))
	e.GET("/users", handlers.GetUsers(db))

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
