package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1.创建路由
	r := gin.Default()

	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	// r.GET("/hello", func(c *gin.Context) {
	// 	// 第一：解析get/post请求的数据
	// 	c.String(http.StatusOK, "hello World!")
	// })

	//2.1 无参路由
	// r.GET("/hello", HelloHandler)
	//2.2 API路由 ：/book/1
	r.GET("/book/:id", GetBookDetailHandler)
	r.GET("/user", GetUserDetailHandler)

	fmt.Println("http://127.0.0.1:8000/")
	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	r.Run(":8000")
}

// 把handler处理函数拿出来
// func HelloHandler(c *gin.Context) {
// 	c.String(http.StatusOK, "hello World!")
// }

func GetBookDetailHandler(c *gin.Context) {
	bookId := c.Param("id")
	fmt.Println("bookId", bookId)
	c.String(http.StatusOK, "hello bookId!")
}

func GetUserDetailHandler(c *gin.Context) {
	//1.获取值，如果没有为nil
	name := c.Query("name")
	//2.获取值，如果没有，使用默认值
	name2 := c.DefaultQuery("name", "default val")
	fmt.Println("获取的用户名", name)
	fmt.Println("获取的用户名", name2)
	c.String(200, "URL params")
}
