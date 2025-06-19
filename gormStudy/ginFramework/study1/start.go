package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	// 调用 run 函数来启动 Gin 框架的 HTTP 服务器
	//run()
	//普通路由
	//normalRouter()

	// 路由分组
	//groupRouter()

	// RESTful API 路由
	//restfulRouter()
	// 重定向
	//redirectRouter()
	// 静态文件
	//staticFileRouter()

	// HTML 模板
	//htmlRouter()
	// 参数绑定
	//bindRouter()
	// 中间件调用栈
	//Middleware()
	//authMiddleware()

	// 验证器
	Validator()

}

// 运行代码：go run start.go
// 访问地址：http://localhost:8080/ping
// 访问结果：{"message":"pong"}
func run() {
	// 启动 Gin 框架的 HTTP 服务器
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run() // 监听并在 0.0.0.0:8080 上启动服务

}

/*
*
普通路由
*/
func normalRouter() {
	router := gin.Default()
	// GET 请求
	router.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	// POST 请求
	router.POST("/submit", func(c *gin.Context) {
		name := c.PostForm("name")
		c.String(200, "Submitted name: %s", name)
	})

	// 启动服务器
	router.Run(":8081")
}

/*
*
路由分组

	// 访问地址：http://localhost:8082/v1/hello
	// 访问地址：http://localhost:8082/v2/hello
	// 访问地址：http://localhost:8082/v1/submit
	// 访问地址：http://localhost:8082/v2/submit
	// 访问结果：{"message":"pong"} 或 "Hello from v1!" 或 "Hello from v2!"
*/
func groupRouter() {
	router := gin.Default()

	// 创建一个路由分组
	v1 := router.Group("/v1")
	{
		v1.GET("/hello", func(c *gin.Context) {
			c.String(200, "Hello from v1!")
		})
		v1.POST("/submit", func(c *gin.Context) {
			name := c.PostForm("name")
			c.String(200, "Submitted name in v1: %s", name)
		})
	}
	// 创建另一个路由分组
	v2 := router.Group("/v2")
	{
		v2.GET("/hello", func(c *gin.Context) {
			c.String(200, "Hello from v2!")
		})
		v2.POST("/submit", func(c *gin.Context) {
			name := c.PostForm("name")
			c.String(200, "Submitted name in v2: %s", name)
		})
	}

	// 启动服务器
	router.Run(":8082")
}

/*
*
RESTFUL API 路由

	// 访问地址：http://localhost:8083/users
	// 访问地址：http://localhost:8083/users/1
*/
func restfulRouter() {
	router := gin.Default()

	// 定义 RESTful API 路由
	router.GET("/users", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Get all users",
		})
	})

	router.POST("/users", func(c *gin.Context) {
		c.JSON(201, gin.H{
			"message": "User created",
		})
	})

	router.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(200, gin.H{
			"message": "Get user with ID: " + id,
		})
	})

	router.PUT("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(200, gin.H{
			"message": "Update user with ID: " + id,
		})
	})

	router.DELETE("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(200, gin.H{
			"message": "Delete user with ID: " + id,
		})
	})

	router.Run(":8083")
}

/*
*
重定向

	// 访问地址：http://localhost:8084/old-route
	// 访问地址：http://localhost:8084/new-route
*/
func redirectRouter() {
	router := gin.Default()

	// 定义重定向路由
	router.GET("/old-route", func(c *gin.Context) {
		c.Redirect(301, "/new-route")
	})

	router.GET("/new-route", func(c *gin.Context) {
		c.String(200, "This is the new route!")
	})

	// 启动服务器
	router.Run(":8084")
}

/*
*
静态文件

	// 访问地址：http://localhost:8085/static/example.txt
	// 访问地址：http://localhost:8085/static-file
*/
func staticFileRouter() {
	router := gin.Default()

	// 设置静态文件目录
	router.Static("/static", "./static")

	// 定义一个路由来访问静态文件
	router.GET("/static-file", func(c *gin.Context) {
		c.File("./static/example.txt") // 假设 static 目录下有 example.txt 文件
	})

	// 启动服务器
	router.Run(":8085")
}

/*
*
HTML
访问：
*/
func htmlRouter() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/**/*")
	router.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
			"title": "Posts",
		})
	})
	router.GET("/users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
			"title": "Users",
		})
	})
	router.Run(":8080")
}

/*
*
参数绑定
*/
func bindRouter() {
	router := gin.Default()

	// 绑定 JSON 参数
	router.POST("/bind/json", func(c *gin.Context) {
		var json struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		}
		if err := c.ShouldBindJSON(&json); err == nil {
			c.JSON(200, gin.H{
				"name": json.Name,
				"age":  json.Age,
			})
		} else {
			c.JSON(400, gin.H{"error": err.Error()})
		}
	})
	// 绑定表单参数
	router.POST("/bind/form", func(c *gin.Context) {
		name := c.PostForm("name")
		age := c.PostForm("age")
		c.JSON(200, gin.H{
			"name": name,
			"age":  age,
		})
	})
	// 绑定查询参数
	//http://localhost:8080/bind/query?name=Tom&age=18
	router.GET("/bind/query", func(c *gin.Context) {

		name := c.Query("name")
		age := c.Query("age")
		c.JSON(200, gin.H{
			"name": name,
			"age":  age,
		})
	})
	// 绑定路径参数
	//http://localhost:8080/bind/path/Tom/18
	router.GET("/bind/path/:name/:age", func(c *gin.Context) {
		name := c.Param("name")
		age := c.Param("age")
		c.JSON(200, gin.H{
			"name": name,
			"age":  age,
		})
	})
	// 启动服务器
	router.Run(":8080")
}

/*
*
调用栈
*/
func mw1() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("mw1 before")
		c.Next()
		fmt.Println("mw1 after")
	}
}
func mw2() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("mw2 before")
		c.Next()
		fmt.Println("mw2 after")
	}
}
func Middleware() {
	r := gin.Default()

	r.GET("/", mw1(), mw2(), func(c *gin.Context) {
		fmt.Println("self")
		c.String(http.StatusOK, "self")
	})

	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}

/*
*
用户认证
*/
func authMiddleware() {
	router := gin.Default()

	// 路由组使用 gin.BasicAuth() 中间件
	// gin.Accounts 是 map[string]string 的一种快捷方式
	authorized := router.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo":    "bar",
		"austin": "1234",
		"lena":   "hello2",
		"manu":   "4321",
	}))

	// /admin/secrets 端点
	// 触发 "localhost:8080/admin/secrets
	authorized.GET("/secrets", func(c *gin.Context) {
		// 获取用户，它是由 BasicAuth 中间件设置的
		user := c.MustGet(gin.AuthUserKey).(string)
		if secret, ok := secrets[user]; ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
		}
	})

	// 监听并在 0.0.0.0:8080 上启动服务
	router.Run(":8080")
}

// 模拟一些私人数据
var secrets = gin.H{
	"foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
	"austin": gin.H{"email": "austin@example.com", "phone": "666"},
	"lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
}

type LoginInfo struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"number"`
	Email    string `json:"email" form:"email" binding:"email"`
}

/*
*
验证器
*/
func Validator() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		login := LoginInfo{}
		err := c.ShouldBind(&login)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, login)
	})

	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
