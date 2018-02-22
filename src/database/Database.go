package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
	"github.com/alandwiprasetyo/rest-api/src/models"
)

var database *gorm.DB

const (
	USERNAME = "root"
	PASSWORD = "root"
	DATABASE = "restapidb"
)

func GetDatabase() *gorm.DB {
	if database == nil {
		database, _ = gorm.Open("mysql", fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", USERNAME, PASSWORD, DATABASE))
	}
	return database
}
func DropTable () {
	GetDatabase().DropTable(&models.News{})
	GetDatabase().DropTable(&models.Topic{}, "news_topics")
}