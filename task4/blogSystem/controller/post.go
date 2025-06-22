package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jioXo/golangStudy/task4/blogSystem/models"
	"github.com/jioXo/golangStudy/task4/blogSystem/utils"
)

/*
*
创建文章
*/
func CreatePost(c *gin.Context) {
	// 从 JWT 中获取用户名
	token, _ := utils.ParseTokenFromRequest(c)
	claims := token.Claims.(jwt.MapClaims)
	userId := claims["userId"].(float64)

	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数绑定失败"})
		return
	}

	// 将 userId 从 float64 转换为 int
	post.UserID = int(userId)

	db := utils.ConnectDB()
	if err := db.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建文章失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "文章创建成功", "post": post})
}

/**
实现文章的读取功能，支持获取所有文章列表
*/

func GetPostList(c *gin.Context) {
	log.Println("系统正在运行：GetPostList被调用")
	db := utils.ConnectDB()

	var titles []string
	if err := db.Model(&models.Post{}).Pluck("title", &titles).Error; err != nil {

		c.JSON(500, gin.H{"error": "获取文章列表失败"})
	}
	c.JSON(200, gin.H{"title": titles})

}

/**
根据标题获取文章详细信息
*/

func GetPostInforByTile(c *gin.Context) {
	db := utils.ConnectDB()

	var post models.Post
	name := c.Query("title")
	if err := db.Where("title= ?", name).First(&post).Error; err != nil {
		c.JSON(404, gin.H{"error": "未找到对应文章"})
		return
	}
	c.JSON(200, post)
}

/**
实现文章的更新功能，只有文章的作者才能更新自己的文章。
*/

func UpdatePost(c *gin.Context) {
	//先鉴权
	// 从 JWT 中获取用户名
	token, err := utils.ParseTokenFromRequest(c)
	if err != nil || token == nil || !token.Valid {
		c.JSON(401, gin.H{"error": "无效的token"})
		return
	}
	claims := token.Claims.(jwt.MapClaims)
	userId := claims["userId"].(float64) // JWT 中数字默认是 float64

	// 获取要更新的文章标题
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数绑定失败"})
		return
	}
	fmt.Println(post)

	//获取一下存的参数
	type input struct {
		Title   string
		Content string
	}
	// 声明变量
	var in input
	in.Title = post.Title
	in.Content = post.Content

	db := utils.ConnectDB()

	//查询文章
	if err := db.Where("title =?", post.Title).First(&post).Error; err != nil {
		c.JSON(404, gin.H{"error": "文章不存在"})
		return
	}

	//检查作者是否为本人
	if post.UserID != int(userId) {
		c.JSON(403, gin.H{"error": "只有作者本人才能修改文章"})
		return
	}
	// 更新文章
	post.Title = in.Title
	post.Content = in.Content
	fmt.Println(post)
	if err := db.Save(&post).Error; err != nil {
		c.JSON(500, gin.H{"error": "更新失败"})
		return
	}
	c.JSON(200, gin.H{"message": "更新成功", "post": post})
}

/*
*
实现文章的删除功能，只有文章的作者才能删除自己的文章。
*/
func DeletePost(c *gin.Context) {
	//先鉴权
	// 从 JWT 中获取用户名
	token, err := utils.ParseTokenFromRequest(c)
	if err != nil || token == nil || !token.Valid {
		c.JSON(401, gin.H{"error": "无效的token"})
		return
	}

	claims := token.Claims.(jwt.MapClaims)
	userId := claims["userId"].(float64) // JWT 中数字默认是 float64
	// 获取要更新的文章标题
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数绑定失败"})
		return
	}
	fmt.Println(post)
	db := utils.ConnectDB()
	//查询文章
	if err := db.Where("title =?", post.Title).First(&post).Error; err != nil {
		c.JSON(404, gin.H{"error": "文章不存在"})
		return
	}
	//检查作者是否为本人
	if post.UserID != int(userId) {
		c.JSON(403, gin.H{"error": "只有作者本人才能删除文章"})
		return
	}
	// 删除文章
	if err := db.Delete(&post).Error; err != nil {
		c.JSON(500, gin.H{"error": "删除失败"})
		return
	}

	c.JSON(200, gin.H{"message": "删除成功"})

}
