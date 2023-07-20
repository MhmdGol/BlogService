package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var db *sql.DB
var err error

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "blog_service"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	//=====================================

	handleRequests()
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/post/create", createNewPost).Methods("POST")

	myRouter.HandleFunc("/category/create", createNewCategory).Methods("POST")
	myRouter.HandleFunc("/category/read", readAllCategories).Methods("GET")
	myRouter.HandleFunc("/category/read/{id}", readACategory).Methods("GET")
	myRouter.HandleFunc("/category/update/{id}", updateCategory).Methods("PUT")
	myRouter.HandleFunc("/category/delete/{id}", deleteCategory).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome!")
}
