package initialize

import (
	"sample-go-service/middlewares"
	"sample-go-service/router"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	// 设置跨域中间件
	Router.Use(middlewares.Cors())
	Router.Use(middlewares.GinLogger(), middlewares.GinRecovery(true))
	// 路由分组
	ApiGroup := Router.Group("/v1/")
	router.UserRouter(ApiGroup)
	router.InitBaseRouter(ApiGroup)
	return Router
}
