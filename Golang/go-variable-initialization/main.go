package main

import "fmt"

var a int = b + 1
var b int = 1

func main() {
	fmt.Println(a) // 2
	fmt.Println(b) // 1
}
