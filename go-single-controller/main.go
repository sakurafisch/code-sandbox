package main

import (
	"net/http"
)

type myHandler struct {
}

func (myhandler *myHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("返回的数据，哈哈哈哈哈"))

}

func main() {
	handler := myHandler{}
	server := http.Server{Addr: "localhost:8090", Handler: &handler}
	server.ListenAndServe()

}
