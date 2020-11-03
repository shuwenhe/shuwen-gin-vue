package model

import "time"

type Category struct {
	ID       uint      `json:"id" gorm:"primary_key"`
	Name     string    `json:"name" gorm:"type:varchar(50);not null;unique"`
	CreateAt time.Time `json:"create_at" gorm:"type:timestamp"`
	UpdateAt time.Time `json:"update_at" gorm:"type:timestamp"`
}
