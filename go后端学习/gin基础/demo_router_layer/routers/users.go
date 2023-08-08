package routers

import "github.com/gin-gonic/gin"

// 分层注册路由
func LoadUsers(e *gin.Engine) {
	e.GET("/user", UserHandler)
}

func UserHandler(c *gin.Context) {
	c.String(200, "用户模块")
}
