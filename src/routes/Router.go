package routes

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func SetRoutes() *gin.Engine {
	router = gin.Default()
	Setup()
	return router
}

func GetRouter() *gin.Engine {
	if router == nil {
		return SetRoutes()
	}
	return router
}

func Setup() {
	NewsRoutes()
	TopicRoutes()
}
