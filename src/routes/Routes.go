package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/alandwiprasetyo/rest-api/src/common"
)

func Setup() *gin.Engine {
	NewsRoutes()
	TopicRoutes()
	return common.GetRouter()
}
