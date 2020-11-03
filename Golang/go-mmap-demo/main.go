package main

import (
	"fmt"
	// "golang.org/x/exp/mmap"
	"os"
)

// func main() {
// 	at, _ := mmap.Open("./tmp.txt")
// 	buff := make([]byte, 2)
// 	_, _ = at.ReadAt(buff, 4)
// 	_ = at.Close()
// 	fmt.Println(string(buff))
// }

func main() {
	f, _ := os.OpenFile("tmp.txt", os.O_CREATE|os.O_RDWR, 0644)
	// _, _ = f.WriteAt([]byte("abcdefg"), 0)

	buff := make([]byte, 2)
	_, _ = f.ReadAt(buff, 4)
	_ = f.Close()
	fmt.Println(string(buff))
}
