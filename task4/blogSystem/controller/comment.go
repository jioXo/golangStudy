package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jioXo/golangStudy/task4/blogSystem/models"
	"github.com/jioXo/golangStudy/task4/blogSystem/utils"
)

/*
*
实现评论的创建功能，已认证的用户可以对文章发表评论。
*/
func CreateComment(c *gin.Context) {
	// 从 JWT 中获取用户名
	token, err := utils.ParseTokenFromRequest(c)
	if err != nil || token == nil || !token.Valid {
		c.JSON(401, gin.H{"error": "无效的token"})
		return
	}
	claims := token.Claims.(jwt.MapClaims)
	userId := claims["userId"].(float64) // JWT 中数字默认是 float64

	var comment models.Comment
	comment.UserID = int(userId)
	comment.PostID = 6
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数绑定失败"})
		return
	}
	db := utils.ConnectDB()

	if err := db.Create(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建评论失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "评论创建成功", "comment": comment})
}

/*
*
实现评论的读取功能，支持获取某篇文章的所有评论列表。
*/
func GetCommentAll(c *gin.Context) {
	log.Println("系统正在运行：GetCommentAll被调用")
	db := utils.ConnectDB()

	var comments []string
	if err := db.Model(&models.Comment{}).Pluck("post_Id", &comments).Error; err != nil {
		c.JSON(500, gin.H{"error": "获取评论列表失败"})
	}
	c.JSON(200, gin.H{"title": comments})
}
