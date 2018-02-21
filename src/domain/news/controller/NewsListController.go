package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/alandwiprasetyo/rest-api/src/domain/news/services"
	"net/http"
)

func ListNews(c *gin.Context) {
	service := services.NewsListService{}
	TopicId := c.Query("topicId")
	if TopicId == "" {
		service.ListNews()
	} else {
		service.ListNewsByTopic(TopicId)
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data": gin.H{
			"news": service.News,
		},
	})
}
func ListNewsByTopic(c *gin.Context) {
	service := services.NewsListService{}
	service.ListNewsByTopic(c.Param("topicId"))

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data": gin.H{
			"news": service.News,
		},
	})
}
