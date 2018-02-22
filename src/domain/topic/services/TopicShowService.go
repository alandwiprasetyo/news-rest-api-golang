package services

import (
	"github.com/alandwiprasetyo/rest-api/src/models"
	database2 "github.com/alandwiprasetyo/rest-api/src/database"
)

type TopicShowService struct {
	models.Response
	Topic models.Topic
}

func (res *TopicShowService) ShowTopic(id string) *TopicShowService {
	database := database2.GetDatabase()
	topic := models.Topic{}

	result := database.Where("id = ?", id).First(&topic)
	if result.RecordNotFound() {
		return res
	}
	if result.Error != nil {
		res.Error = result.Error.Error()
	}

	res.IsFound = true
	res.Topic = topic
	return res
}
