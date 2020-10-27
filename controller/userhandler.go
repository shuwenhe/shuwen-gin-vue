package controller

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/shuwenhe/shuwen-gin-vue/dao"
	"github.com/shuwenhe/shuwen-gin-vue/db"

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
		name = RandomString(10)
		return
	}

	log.Println(name, password, phone)

	if dao.IsPhoneExist(phone) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "User exist!"})
		return
	}

	user := model.User{
		Name:     name,
		Password: password,
		Phone:    phone,
	}
	db.DB.Create(&user)

	ctx.JSON(200, gin.H{
		"msg": "Register success!",
	})
}

func RandomString(n int) string { // HPkPlzdLUX
	var letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	result := make([]byte, n)
	rand.Seed(time.Now().Unix())
	for k := range result {
		result[k] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}
