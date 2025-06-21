package models

type Posts struct {
	ID        int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Title     string `gorm:"not null" json:"title"`
	Content   string `gorm:"not null" json:"content"`
	UserID    int    `gorm:"not null;column:user_id" json:"user_id"`
	User      User   `gorm:"foreignKey:UserID;references:ID" json:"user"`
	CreatedAt string `gorm:"not null" json:"created_at"`
	UpdatedAt string `gorm:"not null" json:"updated_at"`
}
