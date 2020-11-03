package router

import (
	"github.com/gin-gonic/gin"
	"github.com/shuwenhe/shuwen-gin-vue/controller"
	"github.com/shuwenhe/shuwen-gin-vue/middleware"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.CORSMiddleware(), middleware.RecoverMiddleware())

	api := r.Group("v1")
	{
		api.POST("/user/register", controller.Register)
		api.POST("/user/login", controller.Login)
		api.GET("/user/info", middleware.AuthMiddleware(), controller.Info)
	}

	categoryRoutes := r.Group("/category")
	categoryController := controller.NewCategoryController()
	categoryRoutes.POST("/add", categoryController.AddCategory)
	categoryRoutes.GET("/getAll", categoryController.GetCategories)
	categoryRoutes.PUT("/update", categoryController.UpdateCategory)
	categoryRoutes.DELETE("/delete", categoryController.DeleteCategoryByID)

	postRoutes := r.Group("/post")
	postRoutes.Use(middleware.AuthMiddleware())
	postController := controller.NewPostController()
	postRoutes.POST("/add", postController.AddPost)
	postRoutes.GET("/getAll", postController.GetPost)
	postRoutes.PUT("/update", postController.UpdatePost)
	postRoutes.DELETE("/delete", postController.DeletePostByID)

	return r
}
