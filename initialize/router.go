package initialize

import (
	"sample-go-service/middlewares"
	"sample-go-service/router"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	Router.Use(middlewares.GinLogger(), middlewares.GinRecovery(true))
	// 路由分组
	ApiGroup := Router.Group("/v1/")
	router.UserRouter(ApiGroup)
	router.InitBaseRouter(ApiGroup)
	return Router
}
