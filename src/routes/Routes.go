package routes

import (
	"github.com/alandwiprasetyo/rest-api/src/common"
)


func Setup()  {
	NewsRoutes()
	TopicRoutes()
	common.GetRouter().Run(":3000")
}