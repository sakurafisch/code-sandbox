package main

import (
	"fmt"
	"log"
	"net/http"

	"go-orm-demo/user"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", helloWorld).Methods("GET")
	myRouter.HandleFunc("/users", user.AllUsers).Methods("GET")
	myRouter.HandleFunc("/user/{name}/{email}", user.NewUser).Methods("POST")
	myRouter.HandleFunc("/user/{name}", user.DeleteUser).Methods("DELETE")
	myRouter.HandleFunc("/user/{name}/{email}", user.UpdateUser).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	fmt.Println("Program running")

	user.InitialMigration()

	handleRequests()
}
