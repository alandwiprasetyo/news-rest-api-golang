package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/alandwiprasetyo/rest-api/src/domain/topic/services"
	"net/http"
)

func ShowTopic(c *gin.Context) {
	service := services.TopicShowService{}
	service.ShowTopic(c.Param("id"))
	if service.Error != "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": service.Error,
		})
		return
	}

	if service.IsFound {
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
			"data": gin.H{
				"topic": service.Topic,
			},
		})
		return
	}
	c.JSON(http.StatusNotFound, nil)
	return
}
