package migrations

import (
	"github.com/alandwiprasetyo/rest-api/src/models"
	"github.com/alandwiprasetyo/rest-api/src/database"
)

func Migration() {
	db := database.GetDatabase()
	db.CreateTable(&models.News{})
	db.CreateTable(&models.Topic{})

}
