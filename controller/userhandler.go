package controller

import (
	"log"
	"net/http"

	"github.com/shuwenhe/shuwen-gin-vue/response"

	"github.com/shuwenhe/shuwen-gin-vue/dto"

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
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "The phone num must be 11 digits!")
		return
	}
	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "Password not less 6 digits!")
		return
	}
	if len(name) == 0 {
		name = util.RandomString(10)
	}
	if dao.IsPhoneExist(phone) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "User exisit!")
		return
	}
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) // 创建用户的时候要加密用户的密码
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "Hased password error!")
		return
	}
	user := model.User{
		Name:     name,
		Password: string(hasedPassword),
		Phone:    phone,
	}
	db.DB.Create(&user)
	response.Succces(ctx, nil, "Register success!")
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
	response.Succces(ctx, gin.H{"token": token}, "Login success!")
}

func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"user": dto.ToUserDto(user.(model.User))}})
}
