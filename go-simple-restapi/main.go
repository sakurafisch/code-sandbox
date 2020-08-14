package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Article 文章类型
type Article struct {
	Title string `json:"Title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}

// Articles 代表 []Article 切片类型 
type Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func allArticles(w http.ResponseWriter, r *http.Request){
	articles := Articles {
		Article{Title: "Test Title", Desc: "Description", Content: "Hello World"},
	}
	fmt.Println("Endpoint Hit: All Articles Endpoint")
	json.NewEncoder(w).Encode(articles)
}


func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", allArticles)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequests()
}