package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	engine.POST("/register", func(context *gin.Context) {
		fmt.Println(context.FullPath())

		var register Register
		if err := context.ShouldBind(&register); err != nil {
			log.Fatal(err.Error())
			return
		}
		fmt.Println(register.Username)
		fmt.Println(register.Phone)

		context.Writer.Write([]byte(register.Username + " Register "))
	})

	engine.Run()
}

type Register struct {
	Username string `form:"name"`
	Phone    string `form:"phone"`
	Password string `form:"password"`
}
