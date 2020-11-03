package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	engine.GET("/hello", func(context *gin.Context) {
		fmt.Println(context.FullPath())

		name := context.Query("name")
		fmt.Println(name)

		context.Writer.Write([]byte("hello, " + name))
	})

	engine.POST("/login", func(context *gin.Context) {
		fmt.Println(context.FullPath())

		username, exist := context.GetPostForm("username")
		if exist {
			fmt.Println(username)
		}

		if password, exist := context.GetPostForm("password"); exist {
			fmt.Println(password)
		}

		context.Writer.Write([]byte("hello, " + username))
	})

	engine.DELETE("/user/:id", func(context *gin.Context) {
		userID := context.Param("id")
		fmt.Println(userID)
		context.Writer.Write([]byte(" delete 用户ID : " + userID))
	})

	engine.Run(":8080")
}
