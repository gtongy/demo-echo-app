package mysql

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Config struct {
	DataSource string `toml:"dataSource"`
}

func GetDB() *gorm.DB {
	dataSource := dataSource()
	db, err := gorm.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	if err != nil {
		fmt.Println(err.Error())
	}
	return db
}

func dataSource() string {
	address := os.Getenv("CLEARDB_DATABASE_URL")
	if address == "" {
		return defaultDataSource()
	}
	url, err := url.Parse(address)
	if err != nil {
		log.Fatal(err)
	}
	return getProductDataSource(url.User.String(), url.Host, url.Path)
}

func getProductDataSource(user, host, databasePath string) string {
	var dataSourceParams = [...]string{
		user,
		"@tcp(",
		host,
		":3306",
		")",
		databasePath,
		"?parseTime=true",
	}
	var dataSource string
	for _, dataSourceParam := range dataSourceParams {
		dataSource += dataSourceParam
	}
	return dataSource
}

func defaultDataSource() string {
	var config Config
	_, err := toml.DecodeFile("./config.toml", &config)
	if err != nil {
		panic(err)
	}
	return config.DataSource
}
