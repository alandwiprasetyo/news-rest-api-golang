package services

import (
	"github.com/alandwiprasetyo/rest-api/src/database"
	"github.com/alandwiprasetyo/rest-api/src/models/base"
	"github.com/alandwiprasetyo/rest-api/src/models/tables"
)

type TopicListService struct {
	base.Response
	Topic []tables.Topic
}

func (res *TopicListService) ListTopics() *TopicListService {
	database := database.GetDatabase()
	var topic [] tables.Topic

	database.Find(&topic)
	res.Topic = topic
	return res
}
