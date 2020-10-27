package router

import (
	"github.com/gin-gonic/gin"
	"github.com/shuwenhe/shuwen-gin-vue/controller"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("v1")
	{
		api.POST("/user/register", controller.Register)
	}

	return r
}
