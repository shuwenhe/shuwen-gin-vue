package controller

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	name := ctx.PostForm("name") // Get the parameter
	password := ctx.PostForm("password")
	phone := ctx.PostForm("phone")

	if len(phone) != 11 { // Data verification
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "The phone num must be 11 digits!",
		})
		return
	}

	if len(password) < 6 { // verification password
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "Password cannot be less than 6 digits!",
		})
		return
	}

	if len(name) == 0 { // If not input the name, to a 10-digit random string
		name = RandomString(10) // HPkPlzdLUX
		return
	}

	log.Println(name, password, phone) // ShuwenHe 123456 15010729356
	// Create user
	// Return result

	ctx.JSON(200, gin.H{
		"msg": "Register success!",
	})
}

func RandomString(n int) string { // HPkPlzdLUX
	var letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	result := make([]byte, n)
	rand.Seed(time.Now().Unix())
	for k, _ := range result {
		result[k] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}
