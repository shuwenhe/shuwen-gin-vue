package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shuwenhe/utils"
)

func Register(ctx *gin.Context) {
	// name := ctx.PostForm("name") // Get the parameter
	password := ctx.PostForm("password")
	phone := ctx.PostForm("phone")

	if len(phone) != 11 { // Data verification
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "The phone num must be 11 digits!",
		})
	}

	// verification password
	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "Password cannot be less than 6 digits!",
		})
	}

	// Create user

	// Return result

	ctx.JSON(utils.NewSucc("Register success!", gin.H{
		"msg": "Register success!",
	}))
}
