package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Go MySQL Demo")
	db, err := sql.Open("mysql", "root:cmd@tcp(127.0.0.1:3306)/first")
	defer db.Close()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Success to Connect to MySQL")

	insert, err := db.Query("INSERT INTO users VALUES('ELLIOT')")
	defer insert.Close()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Successfully insert into user table")
}
