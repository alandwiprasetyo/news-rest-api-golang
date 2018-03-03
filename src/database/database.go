package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"fmt"
	"os"
	"github.com/alandwiprasetyo/rest-api/src/models"
)

var database *gorm.DB

func GetDatabase() *gorm.DB {
	godotenv.Load()
	if database == nil {
		database, _ = gorm.Open("mysql",
			fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
				os.Getenv("DB_USERNAME"),
				os.Getenv("DB_PASSWORD"),
				os.Getenv("DB_HOST"),
				os.Getenv("DB_PORT"),
				os.Getenv("DB_DATABASE")))
	}
	return database
}
func DropTable() {
	GetDatabase().DropTable(&models.News{})
	GetDatabase().DropTable(&models.Topic{}, "news_topics")
}
