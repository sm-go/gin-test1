package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/test", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusMovedPermanently, "https://www.google.com")
	})

	r.POST("/test", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusFound, "foo")
	})

	r.Run()
}
