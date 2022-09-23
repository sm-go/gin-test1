package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/cookie", func(ctx *gin.Context) {
		cookie, err := ctx.Cookie("gin_cookie")
		if err != nil {
			cookie = "Not Set"
			ctx.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
		}
		fmt.Printf("Cookie value : %s \n", cookie)
	})

	r.Run()
}
