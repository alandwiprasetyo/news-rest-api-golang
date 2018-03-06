package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/alandwiprasetyo/rest-api/src/domain/topic/dto"
	"github.com/alandwiprasetyo/rest-api/src/domain/topic/services"
)

func CreateTopic(c *gin.Context) {
	var json dto.TopicDTO
	c.BindJSON(&json)
	service := services.TopicCreateService{}
	service.CreateTopic(&json)

	if service.ErrorValidation != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":   http.StatusBadRequest,
			"status": "fail",
			"data": gin.H{
				"validation": service.ErrorValidation,
			},
		})
		return
	}

	if service.Error != "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"status":  "error",
			"message": service.Error,
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"data": gin.H{
			"topics": service.Topic,
		},
	})
}
