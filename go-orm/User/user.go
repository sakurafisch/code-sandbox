package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var err error

// User 表的字段
type User struct {
	gorm.Model
	Name  string
	Email string
}

// InitialMigration hhh
func InitialMigration() {
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database")
	}
	if db != nil {
		defer db.Close()
	}
	db.AutoMigrate(&User{})
}

// AllUsers hhhh
func AllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "All Users Endpoint Hit")
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Could not connect to the database")
	}
	if db != nil {
		defer db.Close()
	}
	var users []User
	db.Find(&users)
	json.NewEncoder(w).Encode(users)
}

// NewUser hhhh
func NewUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "New User Endpoint Hit")
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Could not connect to the database")
	}
	if db != nil {
		defer db.Close()
	}
	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]
	db.Create(&User{Name: name, Email: email})
	fmt.Fprintln(w, "New User Successfully Created")
}

// DeleteUser hhhh
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delete User Endpoint Hit")
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Could not connect to the database")
	}
	if db != nil {
		db.Close()
	}
	vars := mux.Vars(r)
	name := vars["name"]
	var user User
	db.Where("name = ?", name).Find(&user)
	db.Delete(&user)
	fmt.Fprintln(w, "Delete User Successfully ")
}

// UpdateUser hhhh
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Update User Endpoint Hit")
	db, err := gorm.Open("sqlite3", test.db)
	if err != nil {
		panic("Could not connect to the database")
	}
	if db != nil {
		db.Close()
	}
	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]
	var user User
	db.Where("name = ?", name).Find(&user)
	user.Email = email
	db.Save(&user)
	fmt.Fprintln(w, "Update User Successfully")
}
