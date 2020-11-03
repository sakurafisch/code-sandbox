package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	engine.Handle("GET", "/hello", func(context *gin.Context) {
		path := context.FullPath()
		fmt.Println(path)

		name := context.DefaultQuery("name", "hello")
		fmt.Println(name)

		context.Writer.Write([]byte("Hello, " + name))
	})

	engine.Handle("POST", "/login", func(context *gin.Context) {
		fmt.Println(context.FullPath)

		username := context.PostForm("username")
		password := context.PostForm("password")
		fmt.Println(username)
		fmt.Println(password)

		context.Writer.Write([]byte(username + "登录"))
	})

	engine.Run(":8090")
}
