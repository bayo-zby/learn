package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	username string `JSON:"username"`
	password string `JSON:"passwd"`
}

/*
 * Login
 * 首页内容
 *
 */
func Login(c *gin.Context) {
	// username := c.Param("user")
	// password := c.Param("passwd")
	c.HTML(http.StatusOK, "user/login.html", nil)
}

func Pass(c *gin.Context) {
	u := &User{}
	err := c.ShouldBind(u)
	if err != nil {
		c.Request.Context().Err()
	}
	c.JSON(200, gin.H{
		"message":  "success",
		"username": u.username,
		"passwd":   u.password,
	})
}
