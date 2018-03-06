package services

import (
	"github.com/alandwiprasetyo/rest-api/src/domain/topic/dto"
	"github.com/alandwiprasetyo/rest-api/src/common"
	"github.com/alandwiprasetyo/rest-api/src/database"
	"github.com/alandwiprasetyo/rest-api/src/models/base"
	"github.com/alandwiprasetyo/rest-api/src/models/tables"
)

type TopicCreateService struct {
	base.Response
	Topic tables.Topic
}

func (res *TopicCreateService) CreateTopic(dto *dto.TopicDTO) *TopicCreateService {
	err := common.GetValidator().Struct(dto)
	if err != nil {
		res.ErrorValidation = common.GetValidationMessage(err)
		return res
	}
	database := database.GetDatabase()
	topic := tables.Topic{Name: dto.Name, Description: dto.Description}
	created := database.Create(&topic)
	if created.Error == nil {
		res.Topic = topic
	} else {
		res.Error = created.Error.Error()
	}
	return res

}
