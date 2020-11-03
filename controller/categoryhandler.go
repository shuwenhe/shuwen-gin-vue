package controller

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/shuwenhe/shuwen-gin-vue/dao"
	"github.com/shuwenhe/shuwen-gin-vue/db"
	"github.com/shuwenhe/shuwen-gin-vue/model"
	"github.com/shuwenhe/shuwen-gin-vue/response"
)

type ICategoryController interface {
	RestController
}

type CategroyController struct { // 结构体实现接口
	DB *gorm.DB
}

func NewCategoryController() ICategoryController {
	db.DB.AutoMigrate(model.Category{})
	return CategroyController{DB: db.DB}
}

func (c CategroyController) AddCategory(ctx *gin.Context) {
	category := &model.Category{
		Name:     ctx.PostForm("name"),
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}
	ctx.Bind(&category)
	if category.Name == "" {
		response.Fail(ctx, "数据验证错误，分类名称必填", nil)
	} else {
		category, err := dao.AddCategory(category)
		if err != nil {
			panic(err)
			return
		}
		response.Succces(ctx, gin.H{"category": category}, "添加分类成功！")
	}
}

func (c CategroyController) GetCategories(ctx *gin.Context) {
	categories := []*model.Category{}
	categories, err := dao.GetCategories()
	if err != nil {
		response.Fail(ctx, "获取分类失败！", nil)
	}
	response.Succces(ctx, gin.H{"categories": categories}, "获取分类成功！")
}

func (c CategroyController) UpdateCategory(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.PostForm("id"))
	category := &model.Category{
		ID:       uint(id),
		Name:     ctx.PostForm("name"),
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}
	ctx.Bind(&category)
	if category.Name == "" {
		response.Fail(ctx, "数据验证错误，分类名称必填", nil)
	} else {
		dao.UpdateCategory(category)
		response.Succces(ctx, gin.H{"category": category}, "修改分类成功")
	}
}

func (c CategroyController) DeleteCategoryByID(ctx *gin.Context) {
	stringID := ctx.PostForm("id")
	id, _ := strconv.Atoi(stringID)
	dao.DeleteCategoryByID(uint(id))
	response.Succces(ctx, nil, "删除分类成功！")
}
