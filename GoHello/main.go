package main

import (
	"fmt"
	"net/http"
)

func welcome(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(res, "<b>Hello Golang Web!</b>")
}

func main() {
	http.HandleFunc("/", welcome)
	http.ListenAndServe("localhost:8081", nil)
}
