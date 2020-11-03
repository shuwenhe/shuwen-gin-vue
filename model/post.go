package model

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/satori/uuid"
)

type Post struct {
	ID         uuid.UUID `json:"id,omitempty" gorm:"type:char(36);primary_key"`
	UserId     uint      `json:"user_id,omitempty" gorm:"not null"`
	CategoryId uint      `json:"category_id,omitempty" gorm:"not null"`
	Category   *Category
	Title      string    `json:"title,omitempty" gorm:"type:varchar(50);not null"`
	HeadImg    string    `json:"head_img,omitempty"`
	Content    string    `json:"content,omitempty" gorm:"type:text;not null"`
	CreatedAt  time.Time `json:"created_at,omitempty" gorm:"type:timestamp"`
	UpdatedAt  time.Time `json:"updated_at,omitempty" gorm:"type:timestamp"`
}

func (post *Post) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("ID", uuid.NewV4())
}
