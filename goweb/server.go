package main

import (
	"learn/goweb/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 加载静态资源
	r.Static("website/static", "./static")

	// 模板
	// 引号包括文件，双引号仅包含目录
	r.LoadHTMLGlob("website/html/**/*")
	// 单独载入模板
	// r.LoadHTMLFiles("website/html/index.html")

	// 载入首页
	r.Any("/", func(c *gin.Context) {
		c.HTML(200, "index/index.html", nil)
	})

	// 载入分组页面
	handler.Router(r)

	r.Run() //:80
}
