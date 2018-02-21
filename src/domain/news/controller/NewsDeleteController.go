package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/alandwiprasetyo/rest-api/src/domain/news/services"
	"net/http"
)

func DeleteNews(c *gin.Context) {
	service := services.NewsDeleteService{}
	response := service.DeleteNews(c.Param("id"))
	if response.Error != "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": service.Error,
		})
		return
	}

	if service.IsFound {
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
			"data":   nil,
		})
		return
	}

	c.JSON(http.StatusNotFound, nil)
	return
}
