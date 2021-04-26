package main

import (
	"learn/golang/goweb/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 加载静态资源
	r.Static("website/static", "./static")

	// 模板
	r.LoadHTMLGlob("website/html/*/*")

	// 载入首页
	r.Any("/", index)

	// 分组页面
	handler.Router(r)

	r.Run(":4242")
}

/*
 *	首页
 */
func index(ctx *gin.Context) {
	name := ctx.DefaultQuery("name", "GUEST")
	ctx.HTML(200, "index.html", name)
}
