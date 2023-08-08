package main

import (
	"demo_router_layer/routers" //导入一个本地的模块
	"fmt"

	"github.com/gin-gonic/gin"
)

/*
  1.初始化项目
  go mod init xxx
  go mod tidy
  go get
  go.mod 文件指定 moudle demo_router_layer

*/

func main() {
	// r本质 *gin.Engine结构体
	r := gin.Default()
	fmt.Println("http://127.0.0.1:8000/")
	//注册路由
	routers.LoadUsers(r)
	routers.LoadBooks(r)
	r.Run(":8000")
}
