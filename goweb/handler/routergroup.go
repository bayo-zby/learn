package handler

import (
	"learn/goweb/handler/user"

	"github.com/gin-gonic/gin"
)

/*
 * RouterGroup
 * 整理各页面的handlefunc
 */
func Router(r *gin.Engine) {
	// user
	users := r.Group("/user")
	{
		users.GET("/login", user.Login)
		users.POST("/pass", user.Pass)
	}
}
