package routers

import "github.com/gin-gonic/gin"

func LoadBooks(r *gin.Engine) {
	r.GET("/book", BookHandler)
}

func BookHandler(c *gin.Context) {
	c.String(200, "书籍")
}
