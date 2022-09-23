package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("local/file", func(ctx *gin.Context) {
		ctx.File("example-file.go")
	})

	var fs http.FileSystem //...
	r.GET("fs/file", func(ctx *gin.Context) {
		ctx.FileFromFS("fs/file.go", fs)
	})
}
