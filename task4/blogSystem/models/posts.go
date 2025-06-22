package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title   string `gorm:"not null" json:"title"`
	Content string `gorm:"not null" json:"content"`
	UserID  int    `gorm:"not null;column:user_id" json:"user_id"`
	User    User   `gorm:"foreignKey:UserID;references:ID" json:"user"`
}
