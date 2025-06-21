package models

type Comment struct {
	ID        int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Comment   string `gorm:"not null" json:"comment"`
	UserID    int    `gorm:"not null;column:user_id" json:"user_id"`
	User      User   `gorm:"foreignKey:UserID;references:ID" json:"user"`
	PostID    int    `gorm:"not null;column:post_id" json:"post_id"`
	Post      Posts  `gorm:"foreignKey:PostID;references:ID" json:"post"`
	CreatedAt string `gorm:"not null" json:"created_at"`
	UpdatedAt string `gorm:"not null" json:"updated_at"`
}
