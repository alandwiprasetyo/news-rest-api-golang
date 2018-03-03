package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/alandwiprasetyo/rest-api/src/domain/news/dto"
	"github.com/alandwiprasetyo/rest-api/src/domain/news/services"
	"net/http"
)

func CreateNews(c *gin.Context) {
	var json dto.NewsDTO
	c.BindJSON(&json)
	service := services.NewsCreateService{}
	service.CreateNews(&json)

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
			"news": service.News,
		},
	})
}