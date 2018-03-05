package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/alandwiprasetyo/rest-api/src/domain/topic/services"
	"net/http"
)

func ListTopic(c *gin.Context) {
	service := services.TopicListService{}
	service.ListTopics()

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data": gin.H{
			"topics": service.Topic,
		},
	})
}
