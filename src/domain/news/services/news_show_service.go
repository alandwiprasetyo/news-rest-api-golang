package services

import (
	"github.com/alandwiprasetyo/rest-api/src/models"
	database2 "github.com/alandwiprasetyo/rest-api/src/database"
)

type NewsShowService struct {
	models.Response
	News models.News
}

func (res *NewsShowService) ShowNews(id string) *NewsShowService {
	database := database2.GetDatabase()
	news := models.News{}

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
