package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/alandwiprasetyo/rest-api/src/domain/topic/dto"
	"github.com/alandwiprasetyo/rest-api/src/domain/topic/services"
	"net/http"
)

func UpdateTopic(c *gin.Context) {
	var json dto.TopicDTO
	c.BindJSON(&json)

	service := services.TopicUpdateService{}
	service.UpdateTopic(c.Param("id"), &json)

	if service.ErrorValidation != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
			"data": gin.H{
				"validation": service.ErrorValidation,
			},
		})
		return
	}

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
