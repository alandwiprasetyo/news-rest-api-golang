package services

import (
	database "github.com/alandwiprasetyo/rest-api/src/database"
	"github.com/alandwiprasetyo/rest-api/src/models/base"
	"github.com/alandwiprasetyo/rest-api/src/models/tables"
)

type TopicDeleteService struct {
	base.Response
	Topic tables.Topic
}

func (res *TopicDeleteService) DeleteTopic(id string) *TopicDeleteService {
	database := database.GetDatabase()
	topic := tables.Topic{}
	result := database.Where("id = ?", id).Delete(&topic)
	if result.Error != nil {
		res.Error = result.Error.Error()
	} else {
		res.IsFound = true;
		res.Topic = topic
	}
	return res
}
