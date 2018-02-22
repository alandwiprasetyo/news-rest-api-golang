package routes

import (
	"github.com/alandwiprasetyo/rest-api/src/domain/topic/controller"
)

func TopicRoutes() {
	router := GetRouter()
	router.GET("/topics", controller.ListTopic)
	router.POST("/topics", controller.CreateTopic)
	router.PUT("/topics/:id", controller.UpdateTopic)
	router.GET("/topics/:id", controller.ShowTopic)
	router.DELETE("/topics/:id", controller.DeleteTopic)
}
