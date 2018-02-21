package migrations

import (
	"github.com/alandwiprasetyo/rest-api/src/models"
	"github.com/alandwiprasetyo/rest-api/src/database"
)

func Migration() {
	database := database.GetDatabase()
	database.CreateTable(&models.News{})
	database.CreateTable(&models.Topic{})
	//database.Model("news_topics").AddForeignKey("topic_id", "topics(id)", "RESTRICT", "RESTRICT")
	//database.Model("news_topics").AddForeignKey("news_id", "news(id)", "RESTRICT", "RESTRICT")

}
