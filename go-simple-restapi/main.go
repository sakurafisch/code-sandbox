package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	mux "github.com/gorilla/mux"
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

func testPostArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HHHH")
	fmt.Fprintf(w, "Test POST endpoint worked")
}


func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)


	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", allArticles).Methods("GET")
	myRouter.HandleFunc("/articles", testPostArticles).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	handleRequests()
}