package main

// import "fmt"

// T is a common struct
type T struct {
	x int
}

func (t T) String() string {
	return "boo"
}

// func main() {

// 	t := T{123}
// 	fmt.Printf("%v\n", t)
// 	fmt.Printf("%#v\n", t)
// }
