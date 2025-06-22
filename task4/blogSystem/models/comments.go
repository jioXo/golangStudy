package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Comment string `gorm:"not null" json:"comment"`
	UserID  int    `gorm:"not null;column:user_id" json:"user_id"`
	User    User   `gorm:"foreignKey:UserID;references:ID" json:"user"`
	PostID  int    `gorm:"not null;column:post_id" json:"post_id"`
	Post    Post   `gorm:"foreignKey:PostID;references:ID" json:"post"`
}
