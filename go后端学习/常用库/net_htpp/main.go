package main

import (
	"fmt"
	"net/http"
)

/*
使用net/http模块，开发一个简单的web server
  1.提供get请求
  2.提供post请求
*/

/*
开发一个web服务器主要步骤
第一步：路由
第二步：处理函数
    - 解析请求的数据
    http.ResponseWriter：本质是一个接口，定义三个方法，返回数据给浏览器
    http.Request:解析url中的数据，返回给浏览器或者请求方
*/

func main() {
	//第一步：路由
	http.HandleFunc("/req/get", dealGetHandler)
	/*
	  dealGetHandler:处理函数
	*/

	fmt.Println("http://127.0.0.1:8005/req/get")
	//第三步：启动服务
	http.ListenAndServe(":8005", nil)
	/*
			  addr:当前server监听的端口号和IP
		    hanDler：处理函数
	*/
}

//处理函数的名字用驼峰命名  xxxHandler函数名

// 第二步：处理函数（处理get请求）
func dealGetHandler(w http.ResponseWriter, r *http.Request) {
	// 1) 解析请求的数据
	query := r.URL.Query() //返回的就是map切片类型
	// 1.1 通过字典下标取路由参数
	if len(query["name"]) > 0 {
		names := query["name"]
		fmt.Println("字典下标的取值", names)
	}
	// 1.2 通过Get方法取值
	names1 := query.Get("name")
	fmt.Println("字典下标的取值", names1)

	fmt.Println(query)

	//直接返回数据
	w.Write([]byte("hello world"))
}
