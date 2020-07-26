package main

import (
	"fmt"
	"net/http"
)

func first(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Hello First")
}

func second(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Hello Second")
}

func main() {
	server := http.Server{Addr: "localhost:8090"}
	http.HandleFunc("/first", first)
	http.HandleFunc("/second", second)
	server.ListenAndServe()
}
