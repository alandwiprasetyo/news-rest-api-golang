package routes

import (
	"github.com/alandwiprasetyo/rest-api/src/domain/news/controller"
	"github.com/alandwiprasetyo/rest-api/src/common"
)

func NewsRoutes() {
	router := common.GetRouter()
	router.GET("/news", controller.ListNews)
	router.GET("/newss/:topicId", controller.ListNewsByTopic)
	router.POST("/news", controller.CreateNews)
	router.PUT("/news/:id", controller.UpdateNews)
	router.GET("/news/:id", controller.ShowNews)
	router.DELETE("/news/:id", controller.DeleteNews)

}
