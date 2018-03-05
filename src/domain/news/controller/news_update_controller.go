package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/alandwiprasetyo/rest-api/src/domain/news/dto"
	"github.com/alandwiprasetyo/rest-api/src/domain/news/services"
	"net/http"
)

func UpdateNews(c *gin.Context) {
	var json dto.NewsDTO
	c.BindJSON(&json)

	service := services.NewsUpdateService{}
	service.UpdateNews(c.Param("id"), &json)

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
				"news": service.News,
			},
		})
		return
	}
	c.JSON(http.StatusNotFound, nil)
	return
}
