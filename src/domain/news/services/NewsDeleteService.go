package services

import (
	"github.com/alandwiprasetyo/rest-api/src/models"
	database2 "github.com/alandwiprasetyo/rest-api/src/database"
)

type NewsDeleteService struct {
	models.Response
	News models.News
}

func (res *NewsDeleteService) DeleteNews(id string) *NewsDeleteService {
	database := database2.GetDatabase()
	news := models.News{}
	result := database.Where("id = ?", id).Delete(&news)
	if result.Error != nil {
		res.Error = result.Error.Error()
	} else {
		res.IsFound = true;
		res.News = news
	}
	return res
}
