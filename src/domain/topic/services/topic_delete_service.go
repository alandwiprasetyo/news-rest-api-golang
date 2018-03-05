package services

import (
	"github.com/alandwiprasetyo/rest-api/src/models"
	database2 "github.com/alandwiprasetyo/rest-api/src/database"
)

type TopicDeleteService struct {
	models.Response
	Topic models.Topic
}

func (res *TopicDeleteService) DeleteTopic(id string) *TopicDeleteService {
	database := database2.GetDatabase()
	topic := models.Topic{}
	result := database.Where("id = ?", id).Delete(&topic)
	if result.Error != nil {
		res.Error = result.Error.Error()
	} else {
		res.IsFound = true;
		res.Topic = topic
	}
	return res
}
