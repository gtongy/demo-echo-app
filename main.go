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
	migrate(db)

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/tasks", handlers.GetTasks(db))

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

func migrate(db *sql.DB) {
	sql := `
	CREATE TABLE IF NOT EXISTS tasks(
		id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
		name VARCHAR(255) NOT NULL
	);
	`

	_, err := db.Exec(sql)
	// Exit if something goes wrong with our SQL statement above
	if err != nil {
		fmt.Println(err)
	}
}
