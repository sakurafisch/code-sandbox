package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	engine.GET("/hello", func(context *gin.Context) {
		fmt.Println(context.FullPath())

		var student Student
		if err := context.ShouldBindQuery(&student); err != nil {
			log.Fatal(err.Error())
			return
		}

		fmt.Println(student.Name)
		fmt.Println(student.Classes)

		context.Writer.Write([]byte("hello, " + student.Name))
	})

	engine.Run()
}

type Student struct {
	Name    string `form:"name"`
	Classes string `form:"classes"`
}
