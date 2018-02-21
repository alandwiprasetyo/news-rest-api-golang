package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var database *gorm.DB

const PATH = "./database.db"

func GetDatabase() *gorm.DB {
	if database == nil {
		database, _ = gorm.Open("sqlite3", PATH)
	}
	return database
}