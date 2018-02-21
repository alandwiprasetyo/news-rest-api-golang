package services

import (
	"github.com/alandwiprasetyo/rest-api/src/models"
	"github.com/alandwiprasetyo/rest-api/src/database"
)

type NewsListService struct {
	models.Response
	News []models.News
}

func (res *NewsListService) ListNews() *NewsListService {
	database := database.GetDatabase()
	var news [] models.News
	database.Preload("Topic").Find(&news)
	res.News = news
	return res
}

func (res *NewsListService) ListNewsByTopic(id string) *NewsListService {
	database := database.GetDatabase()
	var news [] models.News
	var topic models.Topic

	database.Where("id = ?", id).First(&topic)
	database.Preload("Topic").Model(&topic).Related(&news, "News")

	res.News = news
	return res
}
