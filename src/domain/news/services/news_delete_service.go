package services

import (
	"github.com/alandwiprasetyo/rest-api/src/models/tables"
	"github.com/alandwiprasetyo/rest-api/src/models/base"
	database2 "github.com/alandwiprasetyo/rest-api/src/database"
)

type NewsDeleteService struct {
	base.Response
	News tables.News
}

func (res *NewsDeleteService) DeleteNews(id string) *NewsDeleteService {
	database := database2.GetDatabase()
	news := tables.News{}

	result := database.Where("id = ?", id).First(&news)
	news.Status = tables.DELETED
	database.Save(&news).Delete(&news)
	database.Find(&news, id)

	if result.Error != nil {
		res.Error = result.Error.Error()
	} else {
		res.IsFound = true;
		res.News = news
	}
	return res
}
