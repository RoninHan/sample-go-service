package router

import (
	"sample-go-service/controller"
	"sample-go-service/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	{
		UserRouter.POST("login", controller.PasswordLogin)
		UserRouter.POST("list", middlewares.JWTAuth(), middlewares.IsAdminAuth(), controller.GetUserList)
	}
}
