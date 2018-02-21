package services

import (
	"github.com/alandwiprasetyo/rest-api/src/models"
	"github.com/alandwiprasetyo/rest-api/src/domain/news/dto"
	"github.com/alandwiprasetyo/rest-api/src/common"
	database2 "github.com/alandwiprasetyo/rest-api/src/database"
)

type NewsCreateService struct {
	models.Response
	News models.News
}

func (res *NewsCreateService) CreateNews(dto *dto.NewsDTO) *NewsCreateService {
	err := common.GetValidator().Struct(dto)
	if err != nil {
		res.ErrorValidation = common.GetValidationMessage(err)
		return res
	}
	database := database2.GetDatabase()
	news := models.News{Headline: dto.Headline, Title: dto.Title, Description: dto.Description, Tags: dto.Tags, Status: dto.Status}
	created := database.Create(&news)
	if created.Error == nil {
		res.News = news
	} else {
		res.Error = created.Error.Error()
	}
	return res
}
