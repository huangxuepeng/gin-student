package initialize

import (
	"student/middleware"
	"student/routers"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	Router.Use(middleware.Cors())
	ApiGroup := Router.Group("u/v1")
	routers.InitUserRouter(ApiGroup)
	return Router
}
