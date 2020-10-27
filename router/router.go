package router

import (
	"github.com/gin-gonic/gin"
	"github.com/shuwenhe/shuwen-gin-vue/controller"
)

func Run() {
	r := gin.Default()

	r.GET("/api/user/register", controller.Register)

	r.Run()
}
