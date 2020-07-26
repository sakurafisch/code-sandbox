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
	http.HandleFunc("/", welcome)
	server.ListenAndServe()
}
