package services

import (
	"github.com/alandwiprasetyo/rest-api/src/models"
	"github.com/alandwiprasetyo/rest-api/src/domain/topic/dto"
	"github.com/alandwiprasetyo/rest-api/src/common"
	database2 "github.com/alandwiprasetyo/rest-api/src/database"
)

type TopicCreateService struct {
	models.Response
	Topic models.Topic
}

func (res *TopicCreateService) CreateTopic(dto *dto.TopicDTO) *TopicCreateService {
	err := common.GetValidator().Struct(dto)
	if err != nil {
		res.ErrorValidation = common.GetValidationMessage(err)
		return res
	}
	database := database2.GetDatabase()
	topic := models.Topic{Name: dto.Name, Description: dto.Description}
	created := database.Preload("News").Create(&topic)
	if created.Error == nil {
		res.Topic = topic
	} else {
		res.Error = created.Error.Error()
	}
	return res

}
