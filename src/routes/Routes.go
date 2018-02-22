package routes

import (
	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	return GetRouter()
}
