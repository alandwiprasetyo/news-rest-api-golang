package services

import (
	"github.com/alandwiprasetyo/rest-api/src/models"
	"github.com/alandwiprasetyo/rest-api/src/database"
)

type TopicListService struct {
	models.Response
	Topic []models.Topic
}

func (res *TopicListService) ListTopics() *TopicListService {
	database := database.GetDatabase()
	var topic [] models.Topic

	database.Find(&topic)
	res.Topic = topic
	return res
}
