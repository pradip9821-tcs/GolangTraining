package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type JWT struct {
	token string `json:"token"`
}

type Error struct {
	Message string `json:"message"`
}

var db *sql.DB

func main() {

	var err error
	db, err = sql.Open("mysql", "root:Pradip@1234@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected.")
	defer db.Close()

	insert, err := db.Query("INSERT INTO user (email, password)VALUES ( 'test1@gmail.com', '1234' )")

	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()

	results, err := db.Query("select  * from user")

	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var user User

		err = results.Scan(&user.ID, &user.Email, &user.Password)
		if err != nil {
			panic(err.Error())
		}

		log.Println(user)
	}

	router := mux.NewRouter()

	router.HandleFunc("/", homepage).Methods("GET")
	router.HandleFunc("/signup", homepage).Methods("POST")

	log.Println("Server Listening on 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Homepage invoked.")
	fmt.Fprintf(w, "We're in homepage")
}
