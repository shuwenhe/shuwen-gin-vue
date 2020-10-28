package controller

import (
	"log"
	"net/http"

	"github.com/shuwenhe/shuwen-gin-vue/common"
	"github.com/shuwenhe/shuwen-gin-vue/dao"
	"github.com/shuwenhe/shuwen-gin-vue/db"
	"github.com/shuwenhe/shuwen-gin-vue/util"
	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
	"github.com/shuwenhe/shuwen-gin-vue/model"
)

func Register(ctx *gin.Context) {
	name := ctx.PostForm("name")
	password := ctx.PostForm("password")
	phone := ctx.PostForm("phone")
	if len(phone) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "The phone num must be 11 digits!",
		})
		return
	}
	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "Password cannot be less than 6 digits!"})
		return
	}
	if len(name) == 0 {
		name = util.RandomString(10)
	}
	if dao.IsPhoneExist(phone) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "User exist!"})
		return
	}
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) // 创建用户的时候要加密用户的密码
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "加密错误"})
		return
	}
	user := model.User{
		Name:     name,
		Password: string(hasedPassword),
		Phone:    phone,
	}
	db.DB.Create(&user)
	ctx.JSON(200, gin.H{
		"code": 200,
		"msg":  "Register success!",
	})
}

func Login(ctx *gin.Context) {
	phone := ctx.PostForm("phone") // 获取参数
	password := ctx.PostForm("password")
	if phone == "" {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "Phone num not null!"})
		return
	}
	if password == "" {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "Password not null!"})
		return
	}
	if len(phone) != 11 { // 数据验证
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "Phone num must 11 digits!"})
		return
	}
	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "Password len not less 6 bigits!"})
		return
	}
	var user model.User // 判断手机号是否存在
	db.DB.Where("phone = ?", phone).First(&user)
	if user.ID == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "User not exist!"})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil { // 判断密码是否正确
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "password err!"})
	}
	token, err := common.ReleaseToken(user) // 发放token
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "System err!"})
		log.Printf("token generate error:%v", err)
		return
	}
	ctx.JSON(200, gin.H{ // Return result
		"code": 200,
		"data": gin.H{"token": token},
		"msg":  "Login success!",
	})
}

func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"user": user}})
}
