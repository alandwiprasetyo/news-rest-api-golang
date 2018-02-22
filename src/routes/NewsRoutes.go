package routes

import (
	"github.com/alandwiprasetyo/rest-api/src/domain/news/controller"
)

func NewsRoutes() {
	router := GetRouter()

	router.GET("/news", controller.ListNews)
	router.POST("/news", controller.CreateNews)
	router.GET("/news/:id", controller.ShowNews)
	router.PUT("/news/:id", controller.UpdateNews)
	router.DELETE("/news/:id", controller.DeleteNews)
}
