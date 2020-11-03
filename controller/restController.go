package controller

import "github.com/gin-gonic/gin"

type RestController interface {
	AddCategory(ctx *gin.Context)
	UpdateCategory(ctx *gin.Context)
	GetCategories(ctx *gin.Context)
	DeleteCategoryByID(ctx *gin.Context)
}

type RestPostController interface {
	AddPost(ctx *gin.Context)
	UpdatePost(ctx *gin.Context)
	GetPost(ctx *gin.Context)
	DeletePostByID(ctx *gin.Context)
}
