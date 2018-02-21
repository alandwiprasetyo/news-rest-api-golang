package routes

import (
	"github.com/alandwiprasetyo/rest-api/src/domain/news/controller"
	"github.com/alandwiprasetyo/rest-api/src/common"
)

func NewsRoutes() {
	router := common.GetRouter()

	router.GET("/news", controller.ListNews)
	router.POST("/news", controller.CreateNews)
	router.GET("/news/:id", controller.ShowNews)
	router.PUT("/news/:id", controller.UpdateNews)
	router.DELETE("/news/:id", controller.DeleteNews)
	//router.GET("/news/topics/:topic", GetHandler)
	//router.GET("/news/status/:status", GetHandler)
}

