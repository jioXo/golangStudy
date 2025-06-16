package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// 题目1
	//method1()

	// 题目2
	//method2()

}

/*
*
题目1：模型定义
假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
要求 ：
使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
编写Go代码，使用Gorm创建这些模型对应的数据库表。
*/
func method1() {
	db := ConnectDB()
	// 自动迁移，创建表
	err := db.AutoMigrate(&User{}, &Post{}, &Comment{})
	if err != nil {
		panic(err)
	}

	// 插入一些测试数据，确保用户名唯一且不为空
	user := User{Username: "testuser1", Password: "password"}
	db.Create(&user)

	post := Post{Title: "First Post", Content: "This is the content of the first post.", UserID: user.ID}
	db.Create(&post)

	comment := Comment{Content: "Great post!", PostID: post.ID, UserID: user.ID}
	db.Create(&comment)

}

type User struct {
	ID           uint   `gorm:"primaryKey"`
	Username     string `gorm:"unique;not null"`
	Password     string `gorm:"not null"`
	Posts        []Post `gorm:"foreignKey:UserID"`
	ArticleCount int    `gorm:"default:0"` // 文章数量统计字段
}

type Post struct {
	ID            uint      `gorm:"primaryKey"`
	Title         string    `gorm:"not null"`
	Content       string    `gorm:"not null"`
	UserID        uint      `gorm:"not null"`
	Comments      []Comment `gorm:"foreignKey:PostID"`
	CommentStatus string    `gorm:"default:'有评论'"` // 评论状态字段
}

// Post 创建后自动更新用户的文章数量
func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
	return tx.Model(&User{}).Where("id = ?", p.UserID).UpdateColumn("article_count", gorm.Expr("article_count + ?", 1)).Error
}

type Comment struct {
	ID      uint   `gorm:"primaryKey"`
	Content string `gorm:"not null"`
	PostID  uint   `gorm:"not null"`
	UserID  uint   `gorm:"not null"`
}

// Comment 删除后检查评论数量并更新文章评论状态
func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	var count int64
	if err := tx.Model(&Comment{}).Where("post_id = ?", c.PostID).Count(&count).Error; err != nil {
		return err
	}
	if count == 0 {
		return tx.Model(&Post{}).Where("id = ?", c.PostID).Update("comment_status", "无评论").Error
	}
	return nil
}

// ConnectDB 连接数据库
func ConnectDB() *gorm.DB {
	dsn := "root:admin@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

/*
*
题目2：关联查询
基于上述博客系统的模型定义。
要求 ：
编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
编写Go代码，使用Gorm查询评论数量最多的文章信息。
*/
func method2() {
	// 连接数据库
	db := ConnectDB()
	//创建一些数据
	// user1 := User{Username: "testuser1", Password: "password"}
	// if err := db.Create(&user1).Error; err != nil {
	// 	panic(err)
	// }
	// post1 := Post{Title: "First Post", Content: "This is the content of the first post.", UserID: user1.ID}
	// if err := db.Create(&post1).Error; err != nil {
	// 	panic(err)
	// }
	// post2 := Post{Title: "Second Post", Content: "This is the content of the second post.", UserID: user1.ID}
	// if err := db.Create(&post2).Error; err != nil {
	// 	panic(err)
	// }
	// comment1 := Comment{Content: "Great post!", PostID: post1.ID, UserID: user1.ID}
	// if err := db.Create(&comment1).Error; err != nil {
	// 	panic(err)
	// }
	// comment2 := Comment{Content: "Nice article!", PostID: post1.ID, UserID: user1.ID}
	// if err := db.Create(&comment2).Error; err != nil {
	// 	panic(err)
	// }
	// comment3 := Comment{Content: "Interesting read!", PostID: post2.ID, UserID: user1.ID}
	// if err := db.Create(&comment3).Error; err != nil {
	// 	panic(err)
	// }
	// 查询某个用户发布的所有文章及其对应的评论信息
	var user User
	if err := db.Preload("Posts.Comments").First(&user, "username = ?", "testuser1").Error; err != nil {
		panic(err)
	}
	for _, post := range user.Posts {
		println("文章标题:", post.Title)
		println("文章内容:", post.Content)
		for _, comment := range post.Comments {
			println("评论内容:", comment.Content)
		}
	}
	// 查询评论数量最多的文章信息
	var post Post
	if err := db.Model(&Post{}).Select("posts.*, COUNT(comments.id) as comment_count").
		Joins("LEFT JOIN comments ON comments.post_id = posts.id").
		Group("posts.id").Order("comment_count DESC").Limit(1).First(&post).Error; err != nil {
		panic(err)
	}
	println("评论数量最多的文章标题:", post.Title)
	println("评论数量最多的文章内容:", post.Content)
}

/*
*
题目3：钩子函数
继续使用博客系统的模型。
要求 ：
为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
	// 验证钩子函数效果
	var updatedPost Post
	if err := db.First(&updatedPost, post.ID).Error; err != nil {
		panic(err)
	}
	if updatedPost.CommentStatus == "无评论" {
		println("文章评论状态已更新为无评论")
	} else {
		println("文章评论状态仍然存在评论")
	}

	// 验证用户文章数量
	var updatedUser User
	if err := db.First(&updatedUser, user.ID).Error; err != nil {
		panic(err)
	}
	println("用户文章数量统计字段:", updatedUser.ArticleCount)
}
	db.Create(&comment)

	// 触发钩子函数
	if err := db.Save(&post).Error; err != nil {
		panic(err)
	}
	if err := db.Delete(&comment).Error; err != nil {
		panic(err)
	}
	// 验证钩子函数效果
	var updatedPost Post
	if err := db.First(&updatedPost, post.ID).Error; err != nil {
		panic(err)
	}
	if updatedPost.Comments == nil {
		println("文章评论状态已更新为无评论")
	} else {
		println("文章评论状态仍然存在评论")
	}

}
*/
