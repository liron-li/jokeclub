package models

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"goweb/pkg/setting"
)

var DB *gorm.DB

type Model struct {
	gorm.Model
}

func init() {
	var (
		err                                                           error
		dbConnection, dbName, user, password, host, port, tablePrefix string
	)

	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}

	dbConnection = sec.Key("DB_CONNECTION").String()
	dbName = sec.Key("DB_DATABASE").String()
	user = sec.Key("DB_USERNAME").String()
	password = sec.Key("DB_PASSWORD").String()
	host = sec.Key("DB_HOST").String()
	port = sec.Key("DB_PORT").String()
	tablePrefix = sec.Key("DB_TABLE_PREFIX").String()

	DB, err = gorm.Open(dbConnection, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		port,
		dbName))

	if err != nil {
		log.Println(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	DB.SingularTable(true)
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)
}

func CloseDB() {
	defer DB.Close()
}
