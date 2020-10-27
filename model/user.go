package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(25);not null",json:"name,omitempty"`
	Password string `gorm:"varchar(100); not null",json:"password,omitempty"`
	Phone    string `gorm:"size:255;not null",json:"phone,omitempty"`
}
