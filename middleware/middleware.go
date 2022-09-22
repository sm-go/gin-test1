package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	//create a router without any middleware by default
	r := gin.New()

	//Global middle
	//Logger middleware will write the logs to gin.DefaultWriter
	r.Use(gin.Logger())

	//Recovery middleware
	r.Use(gin.Recovery())

	// r.GET("/benchmark", myBenchLogger(), benchEndpoint)

	//authorized group
	authorized := r.Group("/")
	authorized.Use(AuthRequired)
	{
		authorized.POST("/login", loginPoint)
		authorized.POST("/submit", submitPoint)
		authorized.POST("/read", readPoint)

		//nested group
		testing := authorized.Group("testing")
		testing.GET("/analytics", analyticsEndpoint)
	}
}

func AuthRequired(c *gin.Context) {
	c.Next()

}
func loginPoint(c *gin.Context) {
	message := "this is a message for testing group route"
	c.String(http.StatusOK, message)
}
func submitPoint(c *gin.Context) {
	message := "this is a message for testing group route"
	c.String(http.StatusOK, message)
}
func readPoint(c *gin.Context) {
	message := "this is a message for testing group route"
	c.String(http.StatusOK, message)
}
func analyticsEndpoint(c *gin.Context) {
	message := "this is a message for testing group route"
	c.String(http.StatusOK, message)
}
