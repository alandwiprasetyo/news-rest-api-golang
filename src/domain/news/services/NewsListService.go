package services

import (
	"github.com/alandwiprasetyo/rest-api/src/models"
	"github.com/alandwiprasetyo/rest-api/src/database"
)

type NewsListService struct {
	models.Response
	News []models.News
}

func (res *NewsListService) ListNews(topicId string, status string) *NewsListService {
	database := database.GetDatabase()
	var news [] models.News
	var topic models.Topic
	if status != "" && topicId != "" {
		database.First(&topic,topicId)
		database.Unscoped().Preload("Topic").Model(&topic).Where("status = ?", status).Related(&news, "News")
	} else if status != "" {
		database.Unscoped().Preload("Topic").Where("status = ?", status).Find(&news)
	} else if topicId != "" {
		database.First(&topic, topicId)
		database.Preload("Topic").Model(&topic).Related(&news, "News")
	} else {
		database.Preload("Topic").Find(&news)
	}
	res.News = news
	return res
}

func (res *NewsListService) ListNewsByTopic(id string) *NewsListService {
	database := database.GetDatabase()
	var news [] models.News
	var topic models.Topic

	database.First(&topic, id)
	database.Preload("Topic").Model(&topic).Related(&news, "News")

	res.News = news
	return res
}

func (res *NewsListService) ListNewsByStatus(status string) *NewsListService {
	database := database.GetDatabase()
	var news [] models.News
	database.Preload("Topic").Where("status = ?", status).First(&news)
	res.News = news
	return res
}
