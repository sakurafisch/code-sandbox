package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	engine.LoadHTMLGlob("./html/*")

	engine.Static("/image", "./img")

	engine.GET("/hellohtml", func(context *gin.Context) {
		fullPath := "请求路径：" + context.FullPath()
		fmt.Println(fullPath)

		context.HTML(http.StatusOK, "index.html", gin.H{
			"title":    "GinPage",
			"fullPath": fullPath,
		})
	})

	engine.Run()
}
