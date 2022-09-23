package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginForm struct {
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var form LoginForm
	if c.ShouldBind(&form) == nil {
		if form.User == "user" && form.Password == "password" {
			c.JSON(http.StatusOK, gin.H{
				"status:": "you are logged in",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status:": "unauthorized",
			})
		}
	}
}

func main() {
	r := gin.Default()
	r.POST("/login", Login)
	r.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"msg": "hello"})
	})
	r.Run()
}
