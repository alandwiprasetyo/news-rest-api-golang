package services

import (
	"github.com/alandwiprasetyo/rest-api/src/models"
	"github.com/alandwiprasetyo/rest-api/src/database"
	"github.com/alandwiprasetyo/rest-api/src/common"
	"github.com/alandwiprasetyo/rest-api/src/domain/topic/dto"
)

type TopicUpdateService struct {
	models.Response
	Topic models.Topic
}

func (res *TopicUpdateService) UpdateTopic(id string, dto *dto.TopicDTO) *TopicUpdateService {
	err := common.GetValidator().Struct(dto)
	if err != nil {
		res.ErrorValidation = common.GetValidationMessage(err)
		return res
	}
	topic := models.Topic{}
	result := database.GetDatabase().Preload("Topic").Where("id = ?", id).First(&topic)

	if result.RecordNotFound() {
		return res
	}
	topic.Name = dto.Name
	topic.Description = dto.Description
	updated := database.GetDatabase().Save(&topic)

	if updated.Error == nil {
		res.IsFound = true
		res.Topic = topic
	} else {
		res.Error = updated.Error.Error()
	}

	return res
}
