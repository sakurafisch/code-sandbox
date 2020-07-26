package main

import (
	"html/template"
	"net/http"
)

func welcome(res http.ResponseWriter, req *http.Request) {
	templ, _ := template.ParseFiles("./view/index.html")
	templ.Execute(res, nil)
}

func main() {
	server := http.Server{Addr: ":8090"}

	// 若 url 以 /static 开头，把请求转发给指定的路径
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", welcome)
	server.ListenAndServe()
}
