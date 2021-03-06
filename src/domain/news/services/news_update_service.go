package services

import (
	"github.com/alandwiprasetyo/rest-api/src/database"
	"github.com/alandwiprasetyo/rest-api/src/common"
	"github.com/alandwiprasetyo/rest-api/src/domain/news/dto"
	"github.com/alandwiprasetyo/rest-api/src/models/base"
	"github.com/alandwiprasetyo/rest-api/src/models/tables"
)

type NewsUpdateService struct {
	base.Response
	News tables.News
}

func (res *NewsUpdateService) UpdateNews(id string, dto *dto.NewsDTO) *NewsUpdateService {
	err := common.GetValidator().Struct(dto)
	if err != nil {
		res.ErrorValidation = common.GetValidationMessage(err)
		return res
	}
	news := tables.News{}
	result := database.GetDatabase().Preload("Topic").Where("id = ?", id).First(&news)

	if result.RecordNotFound() {
		return res
	}
	news.Headline = dto.Headline
	news.Title = dto.Title
	news.Description = dto.Description
	news.Tags = dto.Tags
	news.Status = dto.Status

	updated := database.GetDatabase().Save(&news)

	if updated.Error == nil {
		res.IsFound = true
		res.News = news
	} else {
		res.Error = updated.Error.Error()
	}

	return res
}
