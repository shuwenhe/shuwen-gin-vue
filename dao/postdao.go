package dao

import (
	"github.com/shuwenhe/shuwen-gin-vue/db"
	"github.com/shuwenhe/shuwen-gin-vue/model"
)

func AddPost(post *model.Post) (*model.Post, error) {
	err := db.DB.Create(&post).Error
	if err != nil {
		return nil, err
	}
	return post, nil
}

func GetPost() ([]*model.Post, error) {
	posts := []*model.Post{}
	err := db.DB.Find(&posts).Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func UpdatePost(post *model.Post) error {
	err := db.DB.Save(&post).Error
	if err != nil {
		return err
	}
	return nil
}

func DeletePostByID() {

}
