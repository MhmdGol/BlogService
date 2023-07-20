package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "blog_service"
)

type Category struct {
	gorm.Model
	Name  string
	Posts []*Post `gorm:"many2many:post_categories;"`
}

type Post struct {
	gorm.Model
	Title             string
	Text              string
	Creation_time     time.Time
	Modification_time time.Time
	Categories        []*Category `gorm:"many2many:post_categories;"`
}

func main() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Open the connection to the database
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	//=====================================

	err = db.AutoMigrate(&Category{}, &Post{})
	if err != nil {
		panic(err)
	}

	handleRequests()
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	// myRouter.HandleFunc("/post/create", createNewPost).Methods("POST")

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
