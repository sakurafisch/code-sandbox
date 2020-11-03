package main

import (
	"fmt"
	"net/http"
	"text/template"
)

// User is the struct of user
type User struct {
	Name string
	Age  int
}

var mapper = func(res http.ResponseWriter, req *http.Request) {
	templ, err := template.ParseFiles("./view/mapper.html")
	if err != nil {
		fmt.Println("Failed to Parse mapper.html")
	}
	m := make(map[string]interface{})
	m["user"] = User{Name: "大佬", Age: 29}
	m["money"] = 998
	templ.Execute(res, m)
}

var greeting = func(res http.ResponseWriter, req *http.Request) {
	templ, err := template.ParseFiles("./view/greeting.html")
	if err != nil {
		fmt.Println("Failed to Parse greeting.html")
	}
	user := User{Name: "winnerwinter", Age: 18}
	templ.Execute(res, user)

}

var welcome = func(res http.ResponseWriter, req *http.Request) {
	templ, err := template.ParseFiles("./view/index.html")
	if err != nil {
		fmt.Println("Failed to Parse index.html")
	}
	templ.Execute(res, "smiling")
}

func main() {
	server := http.Server{Addr: "localhost: 8899"}
	http.HandleFunc("/mapper", mapper)
	http.HandleFunc("/greeting", greeting)
	http.HandleFunc("/", welcome)
	http.HandleFunc("/welcome", welcome)
	server.ListenAndServe()
}
