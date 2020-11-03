package dao

import (
	"github.com/shuwenhe/shuwen-gin-vue/db"
	"github.com/shuwenhe/shuwen-gin-vue/model"
)

func AddCategory(category *model.Category) (*model.Category, error) {
	err := db.DB.Create(&category).Error
	if err != nil {
		return nil, err
	}
	return category, nil
}

func GetCategories() ([]*model.Category, error) {
	categories := []*model.Category{}
	err := db.DB.Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func UpdateCategory(category *model.Category) error {
	err := db.DB.Save(&category).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteCategoryByID(id uint) error {
	category := &model.Category{}
	err := db.DB.Where("id = ?", id).Delete(&category).Error
	if err != nil {
		return err
	}
	return nil
}
