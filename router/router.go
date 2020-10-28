package router

import (
	"github.com/gin-gonic/gin"
	"github.com/shuwenhe/shuwen-gin-vue/controller"
	"github.com/shuwenhe/shuwen-gin-vue/middleware"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("v1")
	{
		api.POST("/user/register", controller.Register)
		api.POST("/user/login", controller.Login)
		api.GET("/user/info", middleware.AuthMiddleware(), controller.Info)
	}

	return r
}
