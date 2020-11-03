package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

func main() {
	var dataChan = make(chan []byte)
	go func(dataChan chan []byte) {
		if data, err := ioutil.ReadFile("tmp.txt"); err == nil {
			time.Sleep(time.Second * 3) // 先睡 3 秒
			dataChan <- data
		} else {
			log.Fatalln(err)
		}
	}(dataChan)
	go fmt.Println(<-dataChan)
	fmt.Println("说不定我先输出了呢")
	time.Sleep(time.Second * 10) // 睡 10 秒
}

// func rf2(path string, ch chan []byte) {
// 	data, err := ioutil.ReadFile(filepath.Base(path))
// 	if err != nil {
// 		log.Fatalln(err)
// 		ch <- []byte{}
// 		return
// 	}
// 	time.Sleep(time.Second * 3)
// 	ch <- data
// }

// func main() {
// 	var ch = make(chan []byte)
// 	go rf2("tmp.txt", ch)
// 	go fmt.Println(<-ch)
// 	fmt.Println("说不定我先输出了")
// 	time.Sleep(time.Second * 10)
// }

// func rf(path string) <-chan []byte {
// 	var ch = make(chan []byte)
// 	go func() {
// 		data, err := ioutil.ReadFile(filepath.Base(path))
// 		if err != nil {
// 			log.Fatal(err)
// 			ch <- []byte{}
// 		}
// 		time.Sleep(time.Second * 3)
// 		ch <- data
// 	}()
// 	return ch
// }

// func main() {
// 	go fmt.Println(<-rf("tmp.txt"))
// 	fmt.Println("说不定我先输出了")
// 	time.Sleep(time.Second * 10)
// }

// func main() {
// 	var tmp = make([]byte, 50, 50)
// 	var err error
// 	go func() {
// 		tmp, err = ioutil.ReadFile(filepath.Base("tmp.txt"))
// 	}()
// 	time.Sleep(time.Duration(1) * time.Microsecond)
// 	if err != nil {
// 		log.Fatal(err)
// 	} else {
// 		fmt.Println(tmp)
// 	}
// }

// fsunc welcome(res http.ResponseWriter, req *http.Request) {
// 	res.Header().Set("Content-Type", "text/html; charset=utf-8")
// 	fmt.Fprintln(res, "<b>Hello Golang Web!</b>")
// }

// func main() {
// 	http.HandleFunc("/", welcome)
// 	http.ListenAndServe("localhost:8081", nil)
// }
