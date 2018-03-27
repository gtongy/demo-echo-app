package mysql

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Config struct {
	DataSource string `toml:"dataSource"`
}

func GetDB() *gorm.DB {
	var config Config
	_, err := toml.DecodeFile("./config.toml", &config)
	if err != nil {
		panic(err)
	}

	dataSource := config.DataSource
	db, err := gorm.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	if err != nil {
		fmt.Println(err.Error())
	}
	return db
}
