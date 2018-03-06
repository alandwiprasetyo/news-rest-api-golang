package services

import (
	"github.com/alandwiprasetyo/rest-api/src/database"
	"github.com/alandwiprasetyo/rest-api/src/models/base"
	"github.com/alandwiprasetyo/rest-api/src/models/tables"
)

type TopicShowService struct {
	base.Response
	Topic tables.Topic
}

func (res *TopicShowService) ShowTopic(id string) *TopicShowService {
	database := database.GetDatabase()
	topic := tables.Topic{}

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
