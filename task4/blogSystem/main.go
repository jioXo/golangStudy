package main

import (
	//导入的是包
	"github.com/gin-gonic/gin"
	"github.com/jioXo/golangStudy/task4/blogSystem/controller"
	"github.com/jioXo/golangStudy/task4/blogSystem/middleware"
	"github.com/jioXo/golangStudy/task4/blogSystem/models"
	"github.com/jioXo/golangStudy/task4/blogSystem/utils"
)

func main() {
	route()
	//createTables()

}

/*
*
路由
*/
func route() {
	r := gin.Default()
	r.POST("/register", controller.Register)
	r.POST("/login", controller.LoginInfo)

	r.GET("/getPostList", controller.GetPostList)
	r.GET("/getPostInforByTile", controller.GetPostInforByTile)
	r.GET("/getCommentAll", controller.GetCommentAll)

	auth := r.Group("/api", middleware.JWTAuthMiddleware())
	auth.POST("/createPost", controller.CreatePost)
	auth.POST("/updatePost", controller.UpdatePost)
	auth.POST("/deletePost", controller.DeletePost)

	auth.POST("/createComment", controller.CreateComment)
	r.Run(":8080")
}

/*
*
根据结构体自动创建表
*/
func createTables() {
	db := utils.ConnectDB()
	// 自动迁移，创建表
	err := db.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})
	if err != nil {
		panic(err)
	}
}
