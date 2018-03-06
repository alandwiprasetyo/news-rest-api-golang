package services

import (
	database2 "github.com/alandwiprasetyo/rest-api/src/database"
	"github.com/alandwiprasetyo/rest-api/src/models/base"
	"github.com/alandwiprasetyo/rest-api/src/models/tables"
)

type NewsShowService struct {
	base.Response
	News tables.News
}

func (res *NewsShowService) ShowNews(id string) *NewsShowService {
	database := database2.GetDatabase()
	news := tables.News{}

	result := database.Preload("Topic").Where("id = ?", id).First(&news)

	if result.RecordNotFound() {
		return res
	}

	if result.Error != nil {
		res.Error = result.Error.Error()
	}

	res.IsFound = true
	res.News = news
	return res
}
