package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/alandwiprasetyo/rest-api/src/domain/news/services"
	"net/http"
)

func ShowNews(c *gin.Context) {
	service := services.NewsShowService{}
	service.ShowNews(c.Param("id"))
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
