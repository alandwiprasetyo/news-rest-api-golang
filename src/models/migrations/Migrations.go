package migrations

import (
	"github.com/alandwiprasetyo/rest-api/src/models"
	"github.com/alandwiprasetyo/rest-api/src/database"
)

func Migration() {
	database := database.GetDatabase()
	database.CreateTable(&models.News{})
	database.CreateTable(&models.Topic{})
	//database.CreateTable(&models.NewsTopic{})
}
