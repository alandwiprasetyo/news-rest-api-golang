package migrations

import (
	"github.com/alandwiprasetyo/rest-api/src/models"
	"github.com/alandwiprasetyo/rest-api/src/database"
)

func Migration() {
	db := database.GetDatabase()
	db.CreateTable(&models.News{})
	db.CreateTable(&models.Topic{})
	//database.Model("news_topics").AddForeignKey("topic_id", "topics(id)", "RESTRICT", "RESTRICT")
	//database.Model("news_topics").AddForeignKey("news_id", "news(id)", "RESTRICT", "RESTRICT")

}
