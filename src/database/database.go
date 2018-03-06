package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
	"github.com/alandwiprasetyo/rest-api/src/models/tables"
)

var database *gorm.DB

const (
	USERNAME = "root"
	PASSWORD = "root"
	HOST     = "db"
	PORT     = "3306"
	DATABASE = "restapidb"
)

func GetDatabase() *gorm.DB {
	//godotenv.Load()
	//if database == nil {
	//	database, _ = gorm.Open("mysql",
	//		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
	//			os.Getenv("DB_USERNAME"),
	//			os.Getenv("DB_PASSWORD"),
	//			os.Getenv("DB_HOST"),
	//			os.Getenv("DB_PORT"),
	//			os.Getenv("DB_DATABASE")))
	//}

	//godotenv.Load()
	if database == nil {
		database, _ = gorm.Open("mysql",
			fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
				"root",
				"root",
				"localhost",
				"3306",
				"restapidb"))
	}
	return database
}

func DropTable() {
	GetDatabase().DropTable(&tables.News{})
	GetDatabase().DropTable(&tables.Topic{}, "news_topics")
}
