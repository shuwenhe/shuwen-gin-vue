package dao

import (
	"testing"

	"github.com/shuwenhe/shuwen-gin-vue/model"
)

func TestAddPost(t *testing.T) {
	post := &model.Post{}
	AddPost(post)
}
