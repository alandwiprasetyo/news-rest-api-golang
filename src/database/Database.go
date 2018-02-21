package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
)

var database *gorm.DB

//const PATH = "./database.db"
const USERNAME = "root"
const PASSWORD = "root"
const DATABASE = "restapidb"

func GetDatabase() *gorm.DB {
	if database == nil {
		database, _ = gorm.Open("mysql", fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", USERNAME, PASSWORD, DATABASE))
		//database, _ = gorm.Open("sqlite3", PATH)
	}
	return database
}
