package main

import (
	"fmt"
	"net/http"
)

var param = func(res http.ResponseWriter, req *http.Request) {
	header := req.Header
	fmt.Fprintln(res, header)

	// 获取get请求的参数
	req.ParseForm()
	fmt.Fprintln(res, req.Form)
	fmt.Fprintln(res, req.Form["name"])
	fmt.Fprintln(res, req.FormValue("name"))
}

func main() {
	server := http.Server{Addr: "localhost: 8090"}
	http.HandleFunc("/param", param)
	server.ListenAndServe()
}
