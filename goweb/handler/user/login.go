package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// 用户结构体
type User struct {
	username string `form:"username" binding:"required"`
	password string `form:"passwd" binding:"required"`
}

var err error

/*
 * Login
 * 登录页内容
 *
 */
func Login(c *gin.Context) {
	// username := c.Param("username")
	// password := c.Param("passwd")
	c.HTML(http.StatusOK, "user/login.html", nil)
}

/**
 * 检查登录信息，并绑定结构体
 */
func Check(c *gin.Context) {
	u := &User{}

	conType := c.Request.Header.Get("Content-Type")
	switch conType {
	case "application/json":
		err = c.BindJSON(u)
	case "application/x-www-form-urlencoded":
		err = c.BindWith(u, binding.Form)
	}

	// 错误处理
	if err := c.ShouldBind(u); err != nil {
		c.JSON(401, gin.H{"status": "unauthorized"})
		return
	}

	// 访问正常
	c.JSON(200, gin.H{
		"message":  "success",
		"username": u.username,
		"passwd":   u.password,
	})
}
