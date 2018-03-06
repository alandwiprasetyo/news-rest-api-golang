package migrations

import (
	"github.com/alandwiprasetyo/rest-api/src/database"
	"github.com/alandwiprasetyo/rest-api/src/models/tables"
)

func Migration() {
	db := database.GetDatabase()
	db.CreateTable(&tables.News{})
	db.CreateTable(&tables.Topic{})
}
