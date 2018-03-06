package routes

import (
	. "github.com/alandwiprasetyo/rest-api/src/domain/news/controller"
	"github.com/gin-gonic/gin"
)

func NewsRoutes() *gin.Engine {
	router := GetRouter()
	router.GET("/news", ListNews)
	router.POST("/news", CreateNews)
	router.GET("/news/:id", ShowNews)
	router.PUT("/news/:id", UpdateNews)
	router.DELETE("/news/:id", DeleteNews)
	return router;
}
