package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello gin")
	r := gin.Default() //this is default route without middleware
	//with middleware
	// r := gin.New()//
	// r := gin.Logger()//
	// r := gin.Recovery()//
	r.GET("/msg", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello Gin Message",
		})
	})

	r.GET("/login", Login)
	r.POST("/register", Register)

	r.GET("/home", homePage)

	r.GET("/query", queryString)               // /query?name=smith&age=23
	r.GET("/path/:name/:age", queryPath)       //path/smith/23
	r.GET("/idp/:someid/*anypara", idnAnyPara) //idp/2233/something
	r.GET("/defaultpara", defaultPara)

	//post form
	r.GET("/posts", defaultPosts)

	//parameter + post form
	r.GET("/news", defaultParaForm)

	//Query Map
	r.POST("/maps", myQueryMap)

	//Group route
	// for Version 1 => v1.0.0
	v1 := r.Group("/v1")
	{
		v1.GET("/products", getAllProducts)
		v1.GET("/products/:id", getSingleProducts)
		v1.POST("/products", CreateProducts)
		v1.PUT("/products/:id", UpdateProducts)
		v1.DELETE("/products/:id", DeleteProducts)
	}
	// for Version 2 => v2.0.0
	v2 := r.Group("/v2")
	{
		v2.GET("/products", getAllProducts)
		v2.GET("/products/:id", getSingleProducts)
		v2.POST("/products", CreateProducts)
		v2.PUT("/products/:id", UpdateProducts)
		v2.DELETE("/products/:id", DeleteProducts)
	}

	//custom log file print
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("endpoint %v %v %v %v\n", absolutePath, handlerName, nuHandlers, httpMethod)
	}

	r.POST("/foo", func(c *gin.Context) {
		c.JSON(http.StatusOK, "foo")
	})

	r.GET("/bar", func(c *gin.Context) {
		c.JSON(http.StatusOK, "bar")
	})

	r.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, "ok")
	})

	//creating log file
	// Disable Console Color, you don't need console color when writing the logs to file.
	gin.DisableConsoleColor()

	// Logging to a file.
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	//XML/JSON/YAML/ProtoBuf rendering

	// gin.H is a shortcut for map[string]interface{}
	r.GET("/someJSON", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	r.GET("/moreJSON", func(c *gin.Context) {
		// You also can use a struct
		var msg struct {
			Name    string `json:"user"`
			Message string
			Number  int
		}
		msg.Name = "Lena"
		msg.Message = "hey"
		msg.Number = 123
		// Note that msg.Name becomes "user" in the JSON
		// Will output  :   {"user": "Lena", "Message": "hey", "Number": 123}
		c.JSON(http.StatusOK, msg)
	})

	r.GET("/someXML", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	r.GET("/someYAML", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	// r.GET("/someProtoBuf", func(c *gin.Context) {
	// 	reps := []int64{int64(1), int64(2)}
	// 	label := "test"
	// 	// The specific definition of protobuf is written in the testdata/protoexample file.
	// 	data := &protoexample.Test{
	// 		Label: &label,
	// 		Reps:  reps,
	// 	}
	// 	// Note that data becomes binary data in the response
	// 	// Will output protoexample.Test protobuf serialized data
	// 	c.ProtoBuf(http.StatusOK, data)
	// })

	// Use the following code if you need to write the logs to file and console at the same time.
	// gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	//upload the file
	// r.MaxMultipartMemory = 8 << 20
	// r.Static("/", "./public")
	// r.POST("/upload", uploadTheFile)

	// r.GET("/someGet", getting)
	// r.POST("/somePost", posting)
	// r.PUT("/somePut", putting)
	// r.DELETE("/someDelete", deleting)
	// r.PATCH("/somePatch", patching)
	// r.HEAD("/someHead", head)
	// r.OPTIONS("/someOptions", options)

	r.Run()
}

func Login(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hey! this is Loign page...",
	})
}

func Register(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hey! this is to register",
	})
}

func queryString(c *gin.Context) {
	name := c.Query("name") //c.Param("name") //c.DefaultQuery("")
	age := c.Query("age")
	c.JSON(200, gin.H{
		"name": name,
		"age":  age,
	})
}

func queryPath(c *gin.Context) {
	name := c.Query("name")
	age := c.Query("age")
	c.JSON(200, gin.H{
		"name": name,
		"age":  age,
	})
}

func homePage(c *gin.Context) {
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println(err.Error())
	}
	c.JSON(200, gin.H{
		"message": string(value),
	})
}

func idnAnyPara(c *gin.Context) {
	someid := c.Param("someid")
	anypara := c.Param("anypara")
	message := someid + " and " + anypara
	c.String(http.StatusOK, message)
}

func defaultPara(c *gin.Context) {
	firstname := c.DefaultQuery("firstname", "Guest")
	lastname := c.Query("lastname")
	c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
}

func defaultPosts(c *gin.Context) {
	message := c.PostForm("message")
	username := c.DefaultPostForm("Jack", "Anonymous")
	c.JSON(http.StatusOK, gin.H{
		"status":   "posted",
		"message":  message,
		"username": username,
	})
}

func defaultParaForm(c *gin.Context) {
	id := c.Query("id")
	page := c.DefaultQuery("page", "0")
	name := c.PostForm("name")
	message := c.PostForm("message")
	confirm := c.DefaultPostForm("confirm", "no confirm")

	fmt.Printf("id : %s, page : %s, name : %s, message : %s, confirm : %s", id, page, name, message, confirm)

}

func myQueryMap(c *gin.Context) {
	ids := c.QueryMap("ids")
	names := c.PostFormMap("names")

	fmt.Printf("ids: %v; names: %v", ids, names)
}

func getAllProducts(c *gin.Context) {
	message := "this is a message for testing group route"
	c.String(http.StatusOK, message)
}

func getSingleProducts(c *gin.Context) {
	message := "this is a message for testing group route"
	c.String(http.StatusOK, message)
}

func CreateProducts(c *gin.Context) {
	message := "this is a message for testing group route"
	c.String(http.StatusOK, message)
}

func UpdateProducts(c *gin.Context) {
	message := "this is a message for testing group route"
	c.String(http.StatusOK, message)
}

func DeleteProducts(c *gin.Context) {
	message := "this is a message for testing group route"
	c.String(http.StatusOK, message)
}
