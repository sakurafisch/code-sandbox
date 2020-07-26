package main

import (
	"net/http"
	"text/template"
	"time"
)

// MyTransfer is my custom time string style
var MyTransfer = func(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

var welcome = func(res http.ResponseWriter, req *http.Request) {
	fm := template.FuncMap{"mt": MyTransfer}
	templ := template.New("index.html").Funcs(fm)
	templ, _ = templ.ParseFiles("view/index.html")
	time5 := time.Date(2020, 7, 26, 21, 25, 34, 0, time.Local)
	templ.Execute(res, time5)
}

func main() {
	server := http.Server{Addr: "localhost: 8899"}
	http.HandleFunc("/", welcome)
	server.ListenAndServe()
}
